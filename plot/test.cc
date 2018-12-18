// test.cc benchmarks that plot is expected to be able to render 2-weeks worth
// of 5 kinds of data at a sample rate of 1 instance per second in less than 1
// second per frame.

#include <unistd.h>

#include <cassert>
#include <chrono>
#include <functional>
#include <iostream>
#include <string>

#include "plot/plot.h"
#include "render/renderer.h"

constexpr int kLines = 5;
constexpr uint64_t kPointsPerLine = 60 * 60 * 24 * 7 * 2;

std::chrono::milliseconds benchmark(const std::function<void()>& action) {
  auto start = std::chrono::high_resolution_clock::now();
  action();
  auto stop = std::chrono::high_resolution_clock::now();
  return std::chrono::duration_cast<std::chrono::milliseconds>(stop - start);
}

int main() {
  using namespace std::chrono_literals;
  plot::Plot plot;

  for (auto i = 0; i < kLines; i++) {
    for (uint64_t j = 0; j < kPointsPerLine; j++) {
      plot.AddData(std::to_string(i), plot::Data{j, static_cast<double>(j)});
    }
  }

  render::Renderer renderer{"test", 900, 500};
  auto millis = benchmark([&]() { renderer.Render(plot); });
  std::cout << "took " << millis.count() << "ms" << std::endl;
  assert(millis < 1000ms);
}

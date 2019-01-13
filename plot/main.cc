#include <chrono>
#include <cstdio>
#include <iostream>
#include <mutex>
#include <string>
#include <thread>

#include "message/message.h"
#include "plot/plot.h"
#include "render/renderer.h"

void Usage() {
  const auto p = [](const auto& l) { std::cerr << l << std::endl; };
  p("");
  p("./plot");
  p("");
  p("plots greenerthumb JSON messages from STDIN.");
  p("");
  p("An example is:");
  p("");
  p("    ./plot");
  p("");
  p("    < {\"Name\": \"Soil\", \"Timestamp\": 0, \"Moisture\": 0.37}");
  p("    < {\"Name\": \"Air\", \"Timestamp\": 3600, \"Temperature\": 84.5}");
  p("    < {\"Name\": \"Soil\", \"Timestamp\": 3600, \"Moisture\": 0.35}");
  p("    < {\"Name\": \"Air\", \"Timestamp\": 7200, \"Temperature\": 82.1}");
  p("");
  p("This will plot 2 lines labelled 'Soil Moisture' and 'Air Temperature'.");
  p("Each will have 2 points. The 'Soil'-line will start at hour 0 and finish");
  p("at hour 1. The 'Air'-line will start at hour 1 and finish at hour 2.");
  p("The entire plot will occupy 2 hours. The range for 'Soil Moisture' will");
  p("be [0.35, 0.37] and the range for 'Air Temperature' will be");
  p("[82.1, 84.5]");
  p("");
}

constexpr auto name = "plot";
constexpr auto width = 900;
constexpr auto height = 500;

void LogError(const std::exception& exception) {
  char buffer[32];
  auto time = std::time(nullptr);
  auto utc = std::gmtime(&time);
  std::strftime(buffer, 32, "%Y-%m-%d %H:%M:%S", utc);
  std::cerr << "ERROR greenerthumb-plot " << buffer << " - " << exception.what()
            << std::endl;
}

int main(int argc, char** argv) {
  if (argc > 1) {
    Usage();
    return 2;
  }

  std::mutex mutex;
  plot::Plot plot;

  std::thread thread([&]() {
    std::string line;
    while (std::getline(std::cin, line)) {
      try {
        std::lock_guard<std::mutex> lock{mutex};
        message::Message message{line};
        for (const auto& field : message.fields) {
          plot.AddData(field.key, plot::Data{message.timestamp, field.value});
        }
      } catch (std::exception& exception) {
        LogError(exception);
      }
    }
  });
  thread.detach();

  try {
    render::Renderer renderer{name, width, height};
    while (renderer.IsRunning()) {
      {
        std::lock_guard<std::mutex> lock{mutex};
        renderer.Render(plot);
      }
      std::this_thread::sleep_for(std::chrono::milliseconds(100));  // 10 Hz.
    }
  } catch (std::exception& exception) {
    LogError(exception);
  }
}

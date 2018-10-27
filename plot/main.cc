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
  p("plots greenerthumb JSON messages from STDIN.");
  p("");
  p("Example:");
  p("");
  p("    ./plot");
  p("");
  p("    < {\"name\": \"Soil\", \"timestamp\": 0, \"Moisture\": 0.37}");
  p("    < {\"name\": \"Air\", \"timestamp\": 3600, \"Temperature\": 84.5, "
    "\"Humidity\": 0.54}");
  p("    < {\"name\": \"Soil\", \"timestamp\": 3600, \"Moisture\": 0.35}");
  p("    < {\"name\": \"Air\", \"timestamp\": 7200, \"Temperature\": 82.1, "
    "\"Humidity\": 0.51}");
  p("");
  p("This will plot 3 lines labelled 'Soil Moisture', 'Air Temperature', and");
  p("'Air Humidity'. Each will have 2 points. The 'Soil'-line will start at");
  p("hour 0 and finish at hour 1. The 'Air'-lines will start at hour 1 and");
  p("finish at hour 2. The entire plot will occupy 2 hours.");
  p("");
}

constexpr auto name = "plot";
constexpr auto width = 900;
constexpr auto height = 500;

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
      } catch (std::exception exception) {
        std::cerr << exception.what() << std::endl;
      }
    }
  });

  try {
    render::Renderer renderer{name, width, height};
    while (renderer.IsRunning()) {
      {
        std::lock_guard<std::mutex> lock{mutex};
        renderer.Render(plot);
      }
      std::this_thread::sleep_for(std::chrono::milliseconds(100));  // 10 Hz.
    }
  } catch (std::exception exception) {
    std::cerr << exception.what() << std::endl;
  }

  std::fclose(stdin);

  thread.join();
}

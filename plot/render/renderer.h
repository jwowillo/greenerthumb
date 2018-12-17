#pragma once

#include <map>
#include <mutex>
#include <random>
#include <string>
#include <utility>
#include <vector>

#include <GL/glew.h>  // Need to import first.

#include <GLFW/glfw3.h>

#include "primitive.h"

namespace render {

// Renderer renders Plots.
class Renderer {
 public:
  Renderer(std::string, int, int);
  ~Renderer();

  Renderer(const Renderer&) = delete;
  Renderer& operator=(const Renderer&) = delete;

  bool IsRunning() const;

  void Render(const plot::Plot& plot);

 private:
  std::mutex mutex_;
  GLFWwindow* window_;
  int width_;
  int height_;

  Area area_;

  std::default_random_engine engine_;

  std::map<std::string, std::tuple<double, double, double>> colors_for_names_;

  std::uniform_real_distribution<double> distribution_{0, 360};
};

}  // namespace render

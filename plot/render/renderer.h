#pragma once

#include <cstddef>
#include <iomanip>
#include <map>
#include <random>
#include <sstream>
#include <string>
#include <utility>

#include <GL/glew.h>  // Need to import first.

#include <GLFW/glfw3.h>

#include "../plot/plot.h"

#include "color.h"
#include "gl.h"
#include "primitive.h"

namespace render {

namespace {

std::string DoubleToString(double x) {
  std::stringstream tmp;
  tmp << std::setprecision(2) << x;
  return tmp.str();
}

void RenderUnit(const Area& area) {
  Region region{0.05, 0.1, 0.75, 0};
  plot::Point point{0.5, 0.5};
  std::string unit{"Hours"};
  RenderString(area, region, point, unit);
}

void RenderHours(const Area& area, double hours) {
  Region region{0.05, 0.2, 0.75, 0.1};

  RenderString(area, region, plot::Point{0, 0}, DoubleToString(0));
  RenderString(area, region, plot::Point{0.1, 0}, DoubleToString(hours / 10));
  RenderString(area, region, plot::Point{0.2, 0},
               DoubleToString(2 * hours / 10));
  RenderString(area, region, plot::Point{0.3, 0},
               DoubleToString(3 * hours / 10));
  RenderString(area, region, plot::Point{0.4, 0},
               DoubleToString(4 * hours / 10));
  RenderString(area, region, plot::Point{0.5, 0},
               DoubleToString(5 * hours / 10));
  RenderString(area, region, plot::Point{0.6, 0},
               DoubleToString(6 * hours / 10));
  RenderString(area, region, plot::Point{0.7, 0},
               DoubleToString(7 * hours / 10));
  RenderString(area, region, plot::Point{0.8, 0},
               DoubleToString(8 * hours / 10));
  RenderString(area, region, plot::Point{0.9, 0},
               DoubleToString(9 * hours / 10));
  RenderString(area, region, plot::Point{1, 0}, DoubleToString(hours));
}

void RenderLabels(
    const Area& area,
    std::map<std::string, std::tuple<double, double, double>> pairs) {
  Region region{0.8, 0.95, 0.95, 0.05};

  plot::Point point{0.1, 1};
  std::set<std::string> labels;
  std::vector<std::tuple<double, double, double>> colors;

  for (const auto& pair : pairs) {
    labels.insert(pair.first);
    colors.push_back(pair.second);
  }

  RenderPoints(area, region, plot::Point{0, 1}, colors);
  RenderStrings(area, region, point, labels);
}

void RenderBorder(const Area& area) {
  Region region{0.05, 0.97, 0.75, 0.2};

  std::vector<plot::Point> inner_lines;
  inner_lines.push_back(plot::Point{0.1, 1});
  inner_lines.push_back(plot::Point{0.1, 0});
  inner_lines.push_back(plot::Point{0.2, 0});
  inner_lines.push_back(plot::Point{0.2, 1});
  inner_lines.push_back(plot::Point{0.3, 1});
  inner_lines.push_back(plot::Point{0.3, 0});
  inner_lines.push_back(plot::Point{0.4, 0});
  inner_lines.push_back(plot::Point{0.4, 1});
  inner_lines.push_back(plot::Point{0.5, 1});
  inner_lines.push_back(plot::Point{0.5, 0});
  inner_lines.push_back(plot::Point{0.6, 0});
  inner_lines.push_back(plot::Point{0.6, 1});
  inner_lines.push_back(plot::Point{0.7, 1});
  inner_lines.push_back(plot::Point{0.7, 0});
  inner_lines.push_back(plot::Point{0.8, 0});
  inner_lines.push_back(plot::Point{0.8, 1});
  inner_lines.push_back(plot::Point{0.9, 1});
  inner_lines.push_back(plot::Point{0.9, 0});
  RenderLines(area, region, inner_lines, {0.4, 0.4, 0.4}, 1);

  std::vector<plot::Point> outer_lines;
  outer_lines.push_back(plot::Point{0, 0});
  outer_lines.push_back(plot::Point{1, 0});
  outer_lines.push_back(plot::Point{1, 1});
  outer_lines.push_back(plot::Point{0, 1});
  outer_lines.push_back(plot::Point{0, 0});
  RenderLines(area, region, outer_lines, {0.02, 0.02, 0.02}, 2);
}

void RenderData(const Area& area, const std::vector<plot::Point>& line,
                std::tuple<double, double, double> color) {
  Region region{0.05, 0.97, 0.75, 0.2};
  RenderLines(area, region, line, color, 2);
}

}  // namespace

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
  GLFWwindow* window_;
  int width_;
  int height_;

  Area area_;

  std::default_random_engine engine_;

  std::map<std::string, std::tuple<double, double, double>> colors_for_names_;

  std::uniform_real_distribution<double> distribution_{0, 360};
};

Renderer::Renderer(std::string name, int width, int height)
    : engine_{10},  // Seed that looked pleasant after some tries.
      area_{width, height},
      window_{MakeWindow(name, width, height)} {
  glfwSetWindowUserPointer(window_, this);
  auto reshape = [](GLFWwindow* window, int width, int height) {
    static_cast<Renderer*>(glfwGetWindowUserPointer(window))->area_ =
        Area{width, height};
    glViewport(0, 0, width, height);
  };
  glfwSetWindowSizeCallback(window_, reshape);
}

Renderer::~Renderer() { glfwTerminate(); }

bool Renderer::IsRunning() const {
  glfwPollEvents();
  return !glfwWindowShouldClose(window_);
}

void Renderer::Render(const plot::Plot& plot) {
  for (const auto& name : plot.LineNames()) {
    if (colors_for_names_.find(name) == colors_for_names_.end()) {
      colors_for_names_[name] = RandomColor(distribution_, engine_);
    }
  }

  glClear(GL_COLOR_BUFFER_BIT);

  RenderUnit(area_);
  RenderHours(area_, plot.Hours());
  RenderLabels(area_, colors_for_names_);
  for (const auto& pair : colors_for_names_) {
    RenderData(area_, plot.LineForName(pair.first), pair.second);
  }
  RenderBorder(area_);

  glFlush();
  glfwSwapBuffers(window_);
}

}  // namespace render

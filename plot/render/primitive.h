#pragma once

#include <set>
#include <string>
#include <utility>
#include <vector>

#include "../plot/plot.h"

namespace render {

// Area defines the area of the window.
struct Area {
  int width;
  int height;
};

// Region of an Area.
struct Region {
  double left;
  double top;
  double right;
  double bottom;
};

// RenderLines as a line strip.
void RenderLines(const Area& area, const Region& region,
                 const std::vector<plot::Point>& lines,
                 const std::tuple<double, double, double>& color, int width);

// RenderPoints in a vertical line.
void RenderPoints(
    const Area& area, const Region& region, const plot::Point& initial_point,
    const std::vector<std::tuple<double, double, double>>& colors);

// RenderString horizontally centered.
void RenderString(const Area& area, const Region& region,
                  const plot::Point& point, const std::string& line);

// RenderStrings in a vertical line.
void RenderStrings(const Area& area, const Region& region,
                   const plot::Point& point,
                   const std::set<std::string>& lines);

}  // namespace render

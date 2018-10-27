#pragma once

#include <cmath>
#include <random>
#include <utility>

namespace render {

namespace {

// ColorFromHsv converts an HSV color with h in range [0, 360), s in range
// [0, 1], and v in range [0, 1] to an RGB value.
//
// Algorithm taken from
// http://dystopiancode.blogspot.com/2012/06/hsv-rgb-conversion-algorithms-in-c.html.
std::tuple<double, double, double> ColorFromHsv(double h, double s,
                                                double v) noexcept {
  double c = v * s;
  double x = c * (1.0 - fabs(fmod(h / 60, 2) - 1));
  double m = v - c;
  if (h >= 0 && h < 60) {
    return {c + m, x + m, m};
  }
  if (h >= 60 && h < 120) {
    return {x + m, c + m, m};
  }
  if (h >= 120 && h < 180) {
    return {m, c + m, x + m};
  }
  if (h >= 180 && h < 240) {
    return {m, x + m, c + m};
  }
  if (h >= 240 && h < 300) {
    return {x + m, m, c + m};
  }
  return {c + m, m, x + m};
}

}  // namespace

// RandomColor from the HSV space with S = 1.0 and V = 0.45;
std::tuple<double, double, double> RandomColor(
    std::uniform_real_distribution<double>& distribution,
    std::default_random_engine& engine) noexcept {
  const auto s = 1.0;
  const auto v = 0.75;
  return ColorFromHsv(distribution(engine), s, v);
}

}  // namespace render

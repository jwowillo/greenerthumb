#pragma once

#include <random>

namespace render {

// RandomColor from the HSV space with S = 1.0 and V = 0.45;
std::tuple<double, double, double> RandomColor(
    std::uniform_real_distribution<double>& distribution,
    std::default_random_engine& engine) noexcept;

}  // namespace render

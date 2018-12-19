#pragma once

#include <cstddef>
#include <limits>
#include <map>
#include <set>
#include <string>
#include <utility>
#include <vector>

namespace plot {

namespace {

constexpr auto kDoubleMin = std::numeric_limits<double>::lowest();
constexpr auto kDoubleMax = std::numeric_limits<double>::max();
constexpr auto kDoubleEpsilon = std::numeric_limits<double>::epsilon();

constexpr auto kLongMin = std::numeric_limits<uint64_t>::lowest();
constexpr auto kLongMax = std::numeric_limits<uint64_t>::max();
constexpr auto kLongEpsilon = std::numeric_limits<uint64_t>::epsilon();

}  // namespace

// Data that can be plotted.
struct Data {
  uint64_t timestamp;
  double value;
};

bool operator==(const Data& lhs, const Data& rhs);
bool operator!=(const Data& lhs, const Data& rhs);
bool operator<(const Data& lhs, const Data& rhs);
bool operator>(const Data& lhs, const Data& rhs);
bool operator<=(const Data& lhs, const Data& rhs);
bool operator>=(const Data& lhs, const Data& rhs);

// Point on a Plot.
struct Point {
  double x;
  double y;
};

// Plots are  collections of lines made of Points.
class Plot {
 public:
  // Hours the Plot covers.
  double Hours() const noexcept;

  // LineNames for all the Plot's lines.
  std::set<std::string> LineNames() const noexcept;

  // AddData to the Plot.
  void AddData(const std::string&, const Data&) noexcept;

  // LineForName returns the line for the name in the Plot.
  std::vector<Point> LineForName(const std::string&) const noexcept;

  // RangeForName returns the range for the line with the name in the Plot.
  std::pair<double, double> RangeForName(const std::string&) const noexcept;

 private:
  std::map<std::string, std::set<Data>> data_;
  std::map<std::string, std::pair<double, double>> ranges_;

  uint64_t low_timestamp_ = kLongMax;
  uint64_t high_timestamp_ = kLongMin;
};

}  // namespace plot

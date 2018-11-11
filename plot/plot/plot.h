#pragma once

#include <cmath>
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

bool operator==(const Data& lhs, const Data& rhs) {
  return lhs.timestamp == rhs.timestamp;
}

bool operator!=(const Data& lhs, const Data& rhs) {
  return !operator==(lhs, rhs);
}

bool operator<(const Data& lhs, const Data& rhs) {
  return lhs.timestamp != rhs.timestamp ? lhs.timestamp < rhs.timestamp
                                        : lhs.value < rhs.value;
}

bool operator>(const Data& lhs, const Data& rhs) { return operator<(rhs, lhs); }

bool operator<=(const Data& lhs, const Data& rhs) {
  return !operator>(lhs, rhs);
}

bool operator>=(const Data& lhs, const Data& rhs) {
  return !operator<(lhs, rhs);
}

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

 private:
  std::map<std::string, std::set<Data>> data_;
  std::map<std::string, std::pair<double, double>> ranges_;

  uint64_t low_timestamp_ = kLongMax;
  uint64_t high_timestamp_ = kLongMin;
};

double Plot::Hours() const noexcept {
  if (data_.size() == 0) {
    return 0;
  }
  return (high_timestamp_ - low_timestamp_) / 3600.0;
}

std::set<std::string> Plot::LineNames() const noexcept {
  std::set<std::string> line_names;
  for (const auto& line : data_) {
    line_names.insert(line.first);
  }
  return line_names;
}

void Plot::AddData(const std::string& name, const Data& data) noexcept {
  if (data_.find(name) == data_.end()) {
    data_[name] = std::set<Data>{};
    ranges_[name] = std::pair<double, double>{kDoubleMax, kDoubleMin};
  }
  if (data_[name].find(data) != data_[name].end()) {
    // Don't re-add the point to data.
    return;
  }
  if (data.value < ranges_[name].first) {
    ranges_[name].first = data.value;
  }
  if (data.value > ranges_[name].second) {
    ranges_[name].second = data.value;
  }
  if (data.timestamp < low_timestamp_) {
    low_timestamp_ = data.timestamp;
  }
  if (data.timestamp > high_timestamp_) {
    high_timestamp_ = data.timestamp;
  }
  data_[name].insert(data);
}

std::vector<Point> Plot::LineForName(const std::string& name) const noexcept {
  std::vector<Point> points;
  if (data_.find(name) == data_.end()) {
    return points;
  }
  for (const auto& data : data_.at(name)) {
    Point point;
    point.x = (data.timestamp - low_timestamp_) /
              ((high_timestamp_ - low_timestamp_) + kDoubleEpsilon);
    const auto range = ranges_.at(name);
    point.y = (data.value - range.first) /
              ((range.second - range.first) + kDoubleEpsilon);
    points.push_back(point);
  }
  return points;
}

}  // namespace plot

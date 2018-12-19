#include "plot.h"

#include <cmath>

namespace plot {

bool operator==(const Data& lhs, const Data& rhs) {
  return lhs.timestamp == rhs.timestamp;
}

bool operator!=(const Data& lhs, const Data& rhs) {
  return !operator==(lhs, rhs);
}

bool operator<(const Data& lhs, const Data& rhs) {
  return lhs.timestamp < rhs.timestamp;
}

bool operator>(const Data& lhs, const Data& rhs) { return operator<(rhs, lhs); }

bool operator<=(const Data& lhs, const Data& rhs) {
  return !operator>(lhs, rhs);
}

bool operator>=(const Data& lhs, const Data& rhs) {
  return !operator<(lhs, rhs);
}

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

std::pair<double, double> Plot::RangeForName(const std::string& name) const
    noexcept {
  return ranges_.at(name);
}

void Plot::AddData(const std::string& name, const Data& data) noexcept {
  if (data_.find(name) == data_.end()) {
    data_[name] = std::set<Data>{};
    ranges_[name] = std::pair<double, double>{kDoubleMax, kDoubleMin};
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

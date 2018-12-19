#pragma once

#include <set>
#include <string>

namespace message {

// Field in a message.
struct Field {
  std::string key;
  double value;
};

bool operator==(const Field& lhs, const Field& rhs);
bool operator!=(const Field& lhs, const Field& rhs);
bool operator<(const Field& lhs, const Field& rhs);
bool operator>(const Field& lhs, const Field& rhs);
bool operator<=(const Field& lhs, const Field& rhs);
bool operator>=(const Field& lhs, const Field& rhs);

// Messages are sent at a timestamp and have Fields.
struct Message {
  // Message constructed from a JSON string representing a Message.
  Message(const std::string&);

  uint64_t timestamp;
  std::set<Field> fields;
};

}  // namespace message

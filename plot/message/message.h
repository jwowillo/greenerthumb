#pragma once

#include <cmath>
#include <cstddef>
#include <limits>
#include <set>
#include <stdexcept>
#include <string>

#include "json.h"

namespace message {

// Field in a message.
struct Field {
  std::string key;
  double value;
};

bool operator==(const Field& lhs, const Field& rhs) {
  return lhs.key == rhs.key &&
         abs(lhs.value - rhs.value) < std::numeric_limits<double>::epsilon();
}

bool operator!=(const Field& lhs, const Field& rhs) {
  return !operator==(lhs, rhs);
}

bool operator<(const Field& lhs, const Field& rhs) {
  return lhs.key != rhs.key ? lhs.key < rhs.key : lhs.value < rhs.value;
}

bool operator>(const Field& lhs, const Field& rhs) {
  return operator<(rhs, lhs);
}

bool operator<=(const Field& lhs, const Field& rhs) {
  return !operator>(lhs, rhs);
}

bool operator>=(const Field& lhs, const Field& rhs) {
  return !operator<(lhs, rhs);
}

// Messages are sent at a timestamp and have Fields.
struct Message {
  // Message constructed from a JSON string representing a Message.
  Message(const std::string&);

  uint64_t timestamp;
  std::set<Field> fields;
};

Message::Message(const std::string& raw) {
  json::JSON object;
  try {
    object = json::JSON::Load(raw);
  } catch (std::invalid_argument) {
    throw std::invalid_argument{"bad JSON string"};
  }

  const auto name_object = object["Name"];
  if (name_object.IsNull()) {
    throw std::invalid_argument{"must pass name"};
  }
  const auto name = name_object.ToString();
  const auto timestamp_object = object["Timestamp"];
  if (timestamp_object.IsNull()) {
    throw std::invalid_argument{"must pass timestamp"};
  }
  timestamp = timestamp_object.ToInt();

  for (const auto& value : object.ObjectRange()) {
    const auto key = value.first;
    // Ignore mandatory keys.
    if (key == "Name" || key == "Timestamp") {
      continue;
    }

    const auto value_object = value.second;
    const auto type = value_object.JSONType();
    if (type != json::JSON::Class::Floating &&
        type != json::JSON::Class::Integral) {
      throw std::invalid_argument{"bad numeric value"};
    }
    fields.insert(Field{name + " " + value.first, value.second.ToFloat()});
  }
}

}  // namespace message

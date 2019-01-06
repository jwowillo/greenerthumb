#include "message.h"

#include <cmath>
#include <cstddef>
#include <limits>
#include <stdexcept>

#include "json.h"

namespace message {

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

Message::Message(const std::string& raw) {
  json::JSON object;
  try {
    object = json::JSON::Load(raw);
  } catch (std::invalid_argument) {
    throw std::invalid_argument{"bad JSON string: " + raw};
  }

  auto header_object = object["Header"];
  if (header_object.IsNull()) {
    throw std::invalid_argument{
        R"(key "Header" is missing from JSON string )" + raw};
  }

  const auto name_object = header_object["Name"];
  if (name_object.IsNull()) {
    throw std::invalid_argument{
        R"(key "Header/Name" is missing from JSON string )" + raw};
  }
  const auto name = name_object.ToString();

  const auto timestamp_object = header_object["Timestamp"];
  if (timestamp_object.IsNull()) {
    throw std::invalid_argument{
        R"(key "Header/Timestamp" is missing from JSON string )" + raw};
  }
  timestamp = timestamp_object.ToInt();

  for (const auto& value : object.ObjectRange()) {
    const auto key = value.first;
    // Ignore mandatory keys.
    if (key == "Header") {
      continue;
    }

    const auto value_object = value.second;
    const auto type = value_object.JSONType();
    if (type == json::JSON::Class::Floating) {
      fields.insert(Field{name + " " + value.first, value.second.ToFloat()});
    } else if (type == json::JSON::Class::Integral) {
      fields.insert(Field{name + " " + value.first,
                          static_cast<double>(value.second.ToInt())});
    } else {
      throw std::invalid_argument{"bad numeric value " +
                                  value.second.ToString()};
    }
  }
}

}  // namespace message

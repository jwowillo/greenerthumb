#include "primitive.h"

#include <GL/glew.h>

#include <GL/freeglut.h>

namespace render {

namespace {

constexpr plot::Point PointInArea(const Area& area, const Region& region,
                                  const plot::Point& point) noexcept {
  double region_left = region.left * area.width;
  double region_top = region.top * area.height;
  double region_right = region.right * area.width;
  double region_bottom = region.bottom * area.height;
  double region_width = region_right - region_left;
  double region_height = region_top - region_bottom;
  double point_x = point.x * region_width;
  double point_y = point.y * region_height;
  return plot::Point{region_left + point_x, region_bottom + point_y};
}

void Render(const std::vector<float>& vertices,
            const std::vector<float>& colors, uint8_t primitive) {
  GLuint vbos[2];
  glGenBuffers(2, vbos);

  GLuint vertex_buffer = vbos[0];
  GLuint color_buffer = vbos[1];

  glBindBuffer(GL_ARRAY_BUFFER, vertex_buffer);
  glBufferData(GL_ARRAY_BUFFER, vertices.size() * sizeof(float), &vertices[0],
               GL_STREAM_DRAW);
  glBindBuffer(GL_ARRAY_BUFFER, color_buffer);
  glBufferData(GL_ARRAY_BUFFER, colors.size() * sizeof(float), &colors[0],
               GL_STREAM_DRAW);

  GLuint vao;
  glGenVertexArrays(1, &vao);
  glBindVertexArray(vao);

  glEnableVertexAttribArray(0);
  glEnableVertexAttribArray(1);

  glBindBuffer(GL_ARRAY_BUFFER, vertex_buffer);
  glVertexAttribPointer(0, 3, GL_FLOAT, GL_FALSE, 0, (GLvoid*)0);

  glBindBuffer(GL_ARRAY_BUFFER, color_buffer);
  glVertexAttribPointer(1, 3, GL_FLOAT, GL_FALSE, 0, (GLvoid*)0);

  glBindVertexArray(vao);

  glDrawArrays(primitive, 0, colors.size() / 3);
}

}  // namespace

// RenderLines as a line strip.
void RenderLines(const Area& area, const Region& region,
                 const std::vector<plot::Point>& lines,
                 const std::tuple<double, double, double>& color, int width) {
  glLineWidth(width);

  std::vector<float> vertices;
  std::vector<float> colors;

  for (const auto& line : lines) {
    const auto converted =
        PointInArea(area, region, plot::Point{line.x, line.y});
    auto x = converted.x;
    auto y = converted.y;
    x /= area.width;
    y /= area.height;
    x = 2 * (x - 0.5);
    y = 2 * (y - 0.5);
    vertices.push_back(x);
    vertices.push_back(y);
    vertices.push_back(0);
    colors.push_back(std::get<0>(color));
    colors.push_back(std::get<1>(color));
    colors.push_back(std::get<2>(color));
  }

  Render(vertices, colors, GL_LINE_STRIP);
}

// RenderPoints in a vertical line.
void RenderPoints(
    const Area& area, const Region& region, const plot::Point& initial_point,
    const std::vector<std::tuple<double, double, double>>& colors) {
  glPointSize(13);

  const auto converted = PointInArea(area, region, initial_point);

  std::vector<float> vertices;
  for (int i = 0; i < colors.size(); i++) {
    auto x = converted.x;
    auto y = converted.y - (i * (13 + 3)) + 5;
    x /= area.width;
    y /= area.height;
    x = 2 * (x - 0.5);
    y = 2 * (y - 0.5);
    vertices.push_back(x);
    vertices.push_back(y);
    vertices.push_back(0);
  }

  std::vector<float> flat_colors;
  for (const auto& color : colors) {
    flat_colors.push_back(static_cast<float>(std::get<0>(color)));
    flat_colors.push_back(static_cast<float>(std::get<1>(color)));
    flat_colors.push_back(static_cast<float>(std::get<2>(color)));
  }

  Render(vertices, flat_colors, GL_POINTS);
}

// RenderString horizontally centered.
void RenderString(const Area& area, const Region& region,
                  const plot::Point& point, const std::string& line) {
  const auto converted = PointInArea(area, region, point);
  auto x = converted.x;
  auto y = converted.y;
  x -= (line.length() * 9) / 2;
  glWindowPos2i(x, y);
  glColor4f(0.02, 0.02, 0.02, 1.0);
  glutBitmapString(GLUT_BITMAP_9_BY_15,
                   reinterpret_cast<const unsigned char*>(line.c_str()));
}

// RenderStrings in a vertical line.
void RenderStrings(const Area& area, const Region& region,
                   const plot::Point& point,
                   const std::set<std::string>& lines) {
  const auto converted = PointInArea(area, region, point);
  const auto x = converted.x;
  const auto y = converted.y;
  glWindowPos2i(x, y);
  glColor4f(0.02, 0.02, 0.02, 1.0);
  for (const auto& line : lines) {
    const auto appended = line + "\n";
    glutBitmapString(GLUT_BITMAP_9_BY_15,
                     reinterpret_cast<const unsigned char*>(appended.c_str()));
  }
}

}  // namespace render

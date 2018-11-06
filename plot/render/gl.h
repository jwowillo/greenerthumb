#pragma once

#include <stdexcept>
#include <string>

#include <GL/glew.h>

#include <GL/freeglut.h>
#include <GLFW/glfw3.h>

#include "SOIL.h"

#include "icon.h"

namespace render {

namespace {

constexpr auto kVertexShader =
    "# version 320 es\n"
    "\n"
    "layout (location = 0) in vec3 VertexPosition;\n"
    "layout (location = 1) in vec3 VertexColor;\n"
    "\n"
    "out vec3 Color;\n"
    "out vec3 Position;\n"
    "\n"
    "void main() {\n"
    "  Color = VertexColor;\n"
    "  Position = VertexPosition;\n"
    "\n"
    "  gl_Position = vec4(VertexPosition, 1.0);\n"
    "}\n";

constexpr auto kFragmentShader =
    "# version 320 es\n"
    "\n"
    "in mediump vec3 Color;\n"
    "layout (location = 0) out mediump vec4 FragColor;\n"
    "\n"
    "void main() {\n"
    "  FragColor = vec4(Color, 1.0);\n"
    "}\n";

}  // namespace

GLFWwindow* MakeWindow(std::string name, int width, int height) {
  if (!glfwInit()) {
    throw std::runtime_error{"couldn't initialize GLFW"};
  }

  GLFWwindow* window =
      glfwCreateWindow(width, height, name.c_str(), nullptr, nullptr);
  if (!window) {
    throw std::runtime_error{"couldn't make window"};
  }

  GLFWimage icons[1];
  icons[0].pixels =
      SOIL_load_image_from_memory(&ICON[0], ICON.size(), &icons[0].width,
                                  &icons[0].height, 0, SOIL_LOAD_RGBA);
  glfwSetWindowIcon(window, 1, icons);
  SOIL_free_image_data(icons[0].pixels);

  glfwMakeContextCurrent(window);

  if (glewInit() != GLEW_OK) {
    throw std::runtime_error{"couldn't initialize GLEW"};
  }

  int argc = 0;
  glutInit(&argc, nullptr);

  GLuint vertex_shader = glCreateShader(GL_VERTEX_SHADER);
  if (!vertex_shader) {
    throw std::runtime_error{"couldn't make vertex-shader"};
  }

  GLuint fragment_shader = glCreateShader(GL_FRAGMENT_SHADER);
  if (!fragment_shader) {
    throw std::runtime_error{"couldn't make fragment-shader"};
  }

  const GLchar* vertex_shaders[] = {kVertexShader};
  const GLchar* fragment_shaders[] = {kFragmentShader};

  glShaderSource(vertex_shader, 1, vertex_shaders, nullptr);
  glShaderSource(fragment_shader, 1, fragment_shaders, nullptr);

  GLint compile_ok;

  glCompileShader(vertex_shader);
  glGetShaderiv(vertex_shader, GL_COMPILE_STATUS, &compile_ok);
  if (!compile_ok) {
    throw std::runtime_error{"couldn't compile vertex-shader"};
  }
  glCompileShader(fragment_shader);
  glGetShaderiv(fragment_shader, GL_COMPILE_STATUS, &compile_ok);
  if (!compile_ok) {
    throw std::runtime_error{"couldn't compile fragment-shader"};
  }

  GLuint program = glCreateProgram();
  if (!program) {
    throw std::runtime_error{"couldn't create program"};
  }

  glAttachShader(program, vertex_shader);
  glAttachShader(program, fragment_shader);

  glLinkProgram(program);

  GLint link_ok;
  glGetProgramiv(program, GL_LINK_STATUS, &link_ok);
  if (!link_ok) {
    throw std::runtime_error{"couldn't link program"};
  }

  glUseProgram(program);

  glDeleteProgram(program);
  glDeleteShader(vertex_shader);
  glDeleteShader(fragment_shader);

  glClearColor(0.98, 0.98, 0.98, 1.0);

  return window;
}

}  // namespace render

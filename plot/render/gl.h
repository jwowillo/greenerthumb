#pragma once

#include <string>

#include <GL/glew.h>  // Need to import first.

#include <GLFW/glfw3.h>

namespace render {

GLFWwindow* MakeWindow(std::string name, int width, int height);

}  // namespace render

cmake_minimum_required(VERSION 3.14.5)

# project name
project(jwt)

# creates the variable EXEC and sets it to hello
set(EXEC a.out)

# set the C++17 standard
set(CMAKE_CXX_STANDARD 17)

set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

set(CMAKE_CXX_COMPILER clang++)
include_directories(third_party/jwt-cpp/include)

# compiles the files defined by SOURCES to generante the executable defined by EXEC
add_executable(${EXEC} ./main.cc)

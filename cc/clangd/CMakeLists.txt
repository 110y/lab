cmake_minimum_required(VERSION 3.14.5)

# project name
project(hello)

# creates the variable EXEC and sets it to hello
set(EXEC hello)

# set the C++17 standard
set(CMAKE_CXX_STANDARD 17)

set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

set(CMAKE_CXX_COMPILER clang++)

include_directories(source)

# puts all .cpp files inside src to the SOURCES variable
# file(GLOB SOURCES foo/**/*.cpp)

# compiles the files defined by SOURCES to generante the executable defined by EXEC
# add_executable(${EXEC} ${SOURCES})
add_executable(${EXEC} ./source/exe/main.cc)

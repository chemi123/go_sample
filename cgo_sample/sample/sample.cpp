#include "sample.hpp"
#include <iostream>

int num() {
  return 1;
}

void hoge() {
  std::string str = "hoge";
  std::cout << str << std::endl;
}

const char* fuga() {
  return "fuga";
}

void showGoString(const char* str) {
  std::cout << str << std::endl;
}

void handler(const std::string& in, std::string& out) {
  out = "response to : " + in;
}

const char* handlerInterface(const char* in) {
  std::string response;
  const std::string request = std::string(in);
  handler(request, response);
  return response.c_str();
}

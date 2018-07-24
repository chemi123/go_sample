#include "sample.hpp"
#include <iostream>

int num() {
  return 1;
}

void hoge() {
  std::string str = "hoge";
  std::cout << str << std::endl;
}

#ifdef __cplusplus
extern "C" {
#endif
  void fuga(const std::string &fuga) {
      std::cout << fuga << std::endl;
  }
#ifdef __cplusplus
}
#endif

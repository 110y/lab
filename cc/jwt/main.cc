#include <iostream>
#include <string>

#include <jwt-cpp/jwt.h>

int main() {
  std::string token = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ4IjoxfQ.jymgBs9NLp-C_sLt815fFZ_S7yMC93xfPYLHv44_6626xOC0P4Zk5GZ9SXd_B2cBj1oKrrXkeIkqTfbATjQWP32DECWgrdZyLFwY09VsUN1qN4BvXrR1R8vJBwfeS8WlVBHx7_3AgjwgtryOya-YkNj-QRKJ7BOMqklX-HZgkiEIn-rW7SY47K9topJlL5oVJwR5OUKJsEoC_TW6ORNI-k2nIylIHOEtHpY80BxFTUV0Kql7Ll5nXNhKpYhA0IUpEZ50K51UKTqQ1PPRCU5ksVCPI_3BKh6HdB13XKx2NGbijuMnlFO-deXkYUjc--yh6iWe0ILXGt3c9lM4qZlHbA";

  auto decoded = jwt::decode(token);
  for (auto &e : decoded.get_payload_claims())
    std::cout << e.first << " = " << e.second.to_json() << std::endl;
}

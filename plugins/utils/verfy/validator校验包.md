# validator校验包

### **字段间比较**
| **标签**          | **描述**                                           |
|------------------|--------------------------------------------------|
| `eqcsfield`      | 当前字段的值等于另一个相对字段的值                  |
| `eqfield`        | 当前字段的值等于另一个字段的值                      |
| `fieldcontains`  | 当前字段必须包含指定的字符                          |
| `fieldexcludes`  | 当前字段必须不包含指定的字符                        |
| `gtcsfield`      | 当前字段的值大于另一个相对字段的值                  |
| `gtecsfield`     | 当前字段的值大于等于另一个相对字段的值              |
| `gtefield`       | 当前字段的值大于等于另一个字段的值                  |
| `gtfield`        | 当前字段的值大于另一个字段的值                      |
| `ltcsfield`      | 当前字段的值小于另一个相对字段的值                  |
| `ltecsfield`     | 当前字段的值小于等于另一个相对字段的值              |
| `ltefield`       | 当前字段的值小于等于另一个字段的值                  |
| `ltfield`        | 当前字段的值小于另一个字段的值                      |
| `necsfield`      | 当前字段的值不等于另一个相对字段的值                |
| `nefield`        | 当前字段的值不等于另一个字段的值                    |

---

### **网络相关**
| **标签**            | **描述**                                             |
|--------------------|----------------------------------------------------|
| `cidr`            | 必须是合法的无类别域间路由（CIDR）                   |
| `cidrv4`          | 必须是合法的 IPv4 CIDR                              |
| `cidrv6`          | 必须是合法的 IPv6 CIDR                              |
| `datauri`         | 必须是合法的 Data URL                               |
| `fqdn`            | 必须是完全限定域名（FQDN）                          |
| `hostname`        | 必须是符合 RFC 952 的主机名                          |
| `hostname_port`   | 必须是主机名加端口格式                               |
| `hostname_rfc1123`| 必须是符合 RFC 1123 的主机名                         |
| `ip`              | 必须是合法的 IP 地址                                |
| `ip4_addr`        | 必须是合法的 IPv4 地址                              |
| `ip6_addr`        | 必须是合法的 IPv6 地址                              |
| `ip_addr`         | 必须是合法的 IP 地址                                |
| `ipv4`            | 必须是合法的 IPv4 地址                              |
| `ipv6`            | 必须是合法的 IPv6 地址                              |
| `mac`             | 必须是合法的 MAC 地址                               |
| `tcp4_addr`       | 必须是合法的 TCP IPv4 地址                          |
| `tcp6_addr`       | 必须是合法的 TCP IPv6 地址                          |
| `tcp_addr`        | 必须是合法的 TCP 地址                               |
| `udp4_addr`       | 必须是合法的 UDP IPv4 地址                          |
| `udp6_addr`       | 必须是合法的 UDP IPv6 地址                          |
| `udp_addr`        | 必须是合法的 UDP 地址                               |
| `unix_addr`       | 必须是合法的 Unix 域套接字地址                       |
| `uri`             | 必须是合法的 URI 字符串                             |
| `url`             | 必须是合法的 URL 字符串                             |
| `http_url`        | 必须是合法的 HTTP URL                               |
| `url_encoded`     | 必须是合法的 URL 编码字符串                          |
| `urn_rfc2141`     | 必须是符合 RFC 2141 的 URN 字符串                    |

---

### **字符串相关**
| **标签**          | **描述**                                           |
|------------------|--------------------------------------------------|
| `alpha`          | 必须仅包含字母                                    |
| `alphanum`       | 必须仅包含字母和数字                              |
| `alphanumunicode`| 必须仅包含字母和数字（支持 Unicode）               |
| `alphaunicode`   | 必须仅包含字母（支持 Unicode）                     |
| `ascii`          | 必须仅包含 ASCII 字符                             |
| `boolean`        | 必须是布尔值                                      |
| `contains`       | 必须包含指定的字符                                |
| `containsany`    | 必须包含任意一个指定字符                          |
| `containsrune`   | 必须包含指定的 Rune                               |
| `endsnotwith`    | 必须不以指定字符结尾                              |
| `endswith`       | 必须以指定字符结尾                                |
| `excludes`       | 必须不包含指定字符                                |
| `excludesall`    | 必须不包含所有指定字符                            |
| `excludesrune`   | 必须不包含指定的 Rune                             |
| `lowercase`      | 必须是小写                                        |
| `multibyte`      | 必须是多字节字符                                  |
| `number`         | 必须是数字                                        |
| `numeric`        | 必须是数值                                        |
| `printascii`     | 必须是可打印的 ASCII 字符                         |
| `startsnotwith`  | 必须不以指定字符开头                              |
| `startswith`     | 必须以指定字符开头                                |
| `uppercase`      | 必须是大写                                        |

---

### **格式相关**
| **标签**           | **描述**                                             |
|-------------------|----------------------------------------------------|
| `base64`         | 必须是合法的 Base64 编码                            |
| `base64url`      | 必须是合法的 Base64URL 编码                         |
| `base64rawurl`   | 必须是合法的 Base64RawURL 编码                      |
| `bic`            | 必须是合法的国际银行代码（ISO 9362）                 |
| `bcp47_language_tag`| 必须是合法的 BCP 47 语言标签                      |
| `btc_addr`       | 必须是合法的比特币地址                              |
| `credit_card`    | 必须是合法的信用卡号                                |
| `email`          | 必须是合法的电子邮件地址                            |
| `eth_addr`       | 必须是合法的以太坊地址                              |
| `hexadecimal`    | 必须是十六进制字符串                                |
| `hexcolor`       | 必须是合法的十六进制颜色代码                        |
| `isbn`           | 必须是合法的国际标准书号（ISBN）                    |
| `jwt`            | 必须是合法的 JSON Web Token                         |
| `latitude`       | 必须是合法的纬度值                                  |
| `longitude`      | 必须是合法的经度值                                  |
| `rgb`            | 必须是合法的 RGB 颜色值                            |
| `ssn`            | 必须是合法的社会安全号码（SSN）                     |
| `uuid`           | 必须是合法的 UUID                                   |

---

### **其他**
| **标签**           | **描述**                                             |
|-------------------|----------------------------------------------------|
| `len`            | 长度必须等于指定值                                  |
| `max`            | 长度或数值最大为指定值                              |
| `min`            | 长度或数值最小为指定值                              |
| `required`       | 必须填写                                             |
| `unique`         | 必须唯一                                             |
| `dir`            | 必须是合法的目录                                     |
| `file`           | 必须是合法的文件                                     |

---

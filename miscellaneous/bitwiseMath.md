## Bitwise Math Tutorial




### Terminology
`binary system`:        base-2
`bit`:                  short for `binary digit`




### Base-2, Base-10 Examples
base-10: 572
> 5 * 10²
> 7 * 10¹
> 2 * 10⁰

base-2: 11010
> 1 * 2⁴    =   16
> 1 * 2³    =    8
> 0 * 2²    =    0
> 1 * 2¹    =    2
> 0 * 2⁰    =    0      
>           =   26




### Bitwise AND [ & ]
Bitwise AND operates on each bit position of the surrounding expressions independently,
according to this rule:
>
>   If both input bits are 1, the resulting output is 1,
>   otherwise the output is 0.
>
> 0 & 0 == 0
> 0 & 1 == 0
> 1 & 0 == 0
> 1 & 1 == 1




#### Common Uses of Bitwise AND [ & ]
`Masking`:                  Select a particular bit (or bits) from an integer value.
If you wanted to access the least significant bit in a variable x, and store the bit in another variable, y:
> int x = 5             // binary 101
> int y = x & 1         // y == 1
> x = 4                 // binary 100
> y = x & 1             // y == 0




### Bitwise OR [ | ]
Bitwise OR operates on each bit position of the surrounding expressions independently, 
according to this rule:
>
>   If either or both input bits are 1, the resulting output is 1,
>   otherwise the output is 0.
>
> 0 | 0 == 0
> 0 | 1 == 1
> 1 | 0 == 1
> 1 | 1 == 1




#### Common Uses of Bitwise OR [ | ]
Bitwise OR is often used to make sure that a given bit is turned on in a given expression.
To copy the bits from a into b, while making sure the lowest bit is set to 1:
> b = a | 1




### Bitwise XOR [ ^ ]
Bitwise Exclusive OR, or Bitwise XOR, operates on each bit position of the surrounding expressions independently, 
according to this rule:
>
>   If exclusively 1 of the input bits are 1, the resulting output is 1,
>   otherwise the output is 0.
>
> 0 ^ 0 == 0
> 0 ^ 1 == 1
> 1 ^ 0 == 1
> 1 ^ 1 == 0




#### Common Uses of Bitwise OR [ | ]
Bitwise XOR is often used toggle some of the bits in an integer expression while leaving others alone.
To toggle the lowest bit in x and store the result in y:
> y = x ^ 1




### Bitwise NOT [ ~ ]
The Bitwise NOT operator is applied to a single operand to its right. 
>
> Bitwise NOT changes each bit to its opposite.
>
> int a = 103       //  0000000001100111
> int b = ~a        //  1111111110011000 = -104
`Two's complement`:         The encoding of positive and negative numbers 
The highest bit is the `sign bit` and we are changing it from positive to negative.




### Bit Shift Operators [ <<, >> ]
There are 2 operators:
- left shift    -   <<
- right shift   -   >>
>
> These operators cause the bits in the left operand to be shifted left or right
> by the number of positions specified by the right operand.
> 
> int a = 5             //              00000101     5 in decimal
> int b = a << 3        //              00101000    40 in decimal
> int c = b >> 3        //              00000101     5 in decimal




#### Common Uses of Bit Shift Operators [ <<, >> ]
`Left-shift Operator`: It multiplies the left operand by 2 raised to the right operand power. 
To generate powers of 2:
> 1 << 0        ==    1         1 * 2⁰
> 1 << 1        ==    2         1 * 2¹
> 1 << 2        ==    4         1 * 2²
> 1 << 3        ==    8         1 * 2³
> 1 << 4        ==   16         1 * 2⁴
`Right-shift Operator`: The exact behavior depends on the type of underlying data because of the `sign bit`.
`Sign extension`:       In a normal integer, the sign bit is copied into lower bits when bit shifted right.
                        Using an unsigned integer gets around this problem.
To divide by powers of 2:
> int x = 1000
> int y = x >> 3    // integer division of 1000 by 8, y = 125




### Assignment Operators
Bitwise AND, OR, left shift, and right shift all have shorthand assignment operators:
> int x = 1             // 0000000000000001
> x <<= 3               // 0000000000001000     left shift 3
> x |= 3                // 0000000000001011     11 is 3 in binary
> x &= 1                // 0000000000000001
> x ^= 4                // 0000000000000101     100 is 4 in binary
> x ^= 4                // 0000000000000001     100 is 4 in binary
There is no shorthand assignment operator for Bitwise NOT.
> x = ~x





### Links:
- https://playground.arduino.cc/Code/BitMath/
- https://www.youtube.com/watch?v=qq64FrA2UXQ
- https://en.wikipedia.org/wiki/Two%27s_complement

⁰
¹
²
³
⁴
⁵
⁶
# ser-IR (Serial + Infra Red)

<p align="center">
<img src="https://github.com/ncaio/serIR/blob/main/img/Neo%20Tech.png" alt="alt text" width="400" class="center"/>
</p>

Ser-IR is an educational abstraction app of Infra Red signals captured from human and Television interaction  to talk about possible cybersecurity risks, such as:

 - **Keylogging**;
 - **Profile Tracking**;
 - **IR payload injection**.

## Getting started

### How?

When you are in front of your TV holding the controller and pressing any button, a signal will travel from you to the television. Am I correct? This project uses a hardware and software combination to allow you to capture and store those signals.

### The Hardware

This project is based on a NEC Infrared Codec Module (ver1.0) (YS-IRTM).

#### Characteristics

 - 38khz InfraRed encoder(transmitter) and decoder (receiver);
 - Infrared Head Extension (scalability); 
 - UART Microcontroller Serial Communication Interface.

### The software

A virtual/real keyboard is or looks like a matrix. I/We have learned how to abstract it using a programming language. In this case, Golang. For example, the next table represents a multidimensional [4][10] array. 

| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| q | w | e | r | t | y | u | i | o | p |
| a | s | d | f | `g` | h | j | k | l | ç |
| z | x | c | v | b | n | m | , | . | ; |

The algorithm used here interacts with the sent signals walking through this array. And every virtual keyboard has an initial/start point. In this case, the starting point is **g** or index [2,4].   

**Example - LG TV KEYBOARD MAPPING**

Just a piece of code from this project. It is this way of mapping a keyboard as a matrix using Golang.

```
keyboard := [4][10]string{{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}, {"q", "w", "e", "r", "t", "y
", "u", "i", "o", "p"}, {"a", "s", "d", "f", "g", "h", "j", "k", "l", "?"}, {"@", "z", "x", "c", "v", "b", "n", "m",
 ",", "."}}

```
### Keylogging

### Profile Tracking

### IR payload injection


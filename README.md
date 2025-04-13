# serIR (Serial + Infrared)

<p align="center">
<img src="https://github.com/ncaio/serIR/blob/main/img/Neo%20Tech.png" alt="alt text" width="400" class="center"/>
</p>

SerIR é uma ferramenta educacional que aborda o assunto segurança da informação, sendo mais específico, a captura de sinais infravermelhos a partir da comunicação humana e uma TV/Smart Tv. Este app nos possibilita discutirmos sobre os seguintes aspectos: (SerIR is an educational app about Infrared signals captured from human and Television interaction. It's to talking about possible cybersecurity risks, such as:)

 - **IR Keylogging**;
 - **IR Tracking and Behavioral Profiling**;
 - **IR payload injection**.

## Getting started

### How?

Quando você está em frente à TV, segurando o controle e pressionando qualquer botão, um sinal viaja de você para a televisão. Estou certo? Este projeto utiliza uma combinação de hardware e software para permitir a captura e o armazenamento desses sinais. (When you are in front of your TV, holding the controller and pressing any button, a signal travels from you to the television. Am I correct? This project uses a hardware and software combination to allow you to capture and store those )signals.

### Hardware

Hardware inventory: 
 - O1 NEC Infrared Codec Module (ver1.0) (YS-IRTM);
 - 02 BT-05 Bluetooth modules;
 - O1 CP2120 Module USB to TTL.
 - LG Model 43LM6300PSB ~ 50/60 Hz 68 W ~ Software version 05.40.45

#### Characteristics

 - 38khz InfraRed encoder(transmitter) and decoder (receiver);
 - Infrared Head Extension (scalability); 
 - UART Microcontroller Serial Communication Interface;
 - Bluetooth communication.

### The software

Um teclado virtual/real é ou se parece com uma matriz, portanto, podemos abstraí-lo usando uma linguagem de programação, neste caso, Golang. Por exemplo, a próxima tabela representa uma matriz multidimensional [4][10]. ( A virtual/real keyboard is or looks like a matrix, so we can abstract it using a programming language, in this case, Golang. For example, the next table represents a multidimensional [4][10] array.)

| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| q | w | e | r | t | y | u | i | o | p |
| a | s | d | f | `g` | h | j | k | l | ç |
| z | x | c | v | b | n | m | , | . | ; |

O algoritmo interage com os sinais infravermelhos enviados, capturados pelo módulo infravermelho, e reproduz os "movimentos". Obs.: Cada teclado virtual possui um ponto inicial/inicial. O ponto inicial deste teste foi **g** ou índice [2,4]. ( The algorithm interacts with the sent IR signals, captured by the IR module, and replays the "movements". PS: Every virtual keyboard has an initial/start point. This test's starting point was **g** or index [2,4].)

**Example - LG TV KEYBOARD MAPPING**

<p align="center">
<img src="https://github.com/ncaio/serIR/blob/main/img/Screenshot_20240909_164601.png" alt="alt text" width="600" class="center"/>
</p>

O código a seguir representa os caracteres básicos observados na última imagem (entrada de TV). (The following code represents the basic characters observed on the last image (TV input).)

```
keyboard := [4][10]string{{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}, {"q", "w", "e", "r", "t", "y
", "u", "i", "o", "p"}, {"a", "s", "d", "f", "g", "h", "j", "k", "l", "?"}, {"@", "z", "x", "c", "v", "b", "n", "m",
 ",", "."}}
```
### IR Keylogging

Neste cenário, falaremos sobre keyloggers infravermelhos. Como sabemos, o keylogger é um aplicativo maligno criado para registrar entradas de teclado. Há muitas maneiras de fazer isso, dependendo de cada cenário. No nosso caso, estamos capturando sinais infravermelhos. (In this scenario, We'll talk about infrared keyloggers. As We know, the keylogger is an evil app created to record keyboard inputs. There are a lot of ways to do it, it depends on each scenario. In our case, we are capturing IR signals.)

Steps:

 - Posicione o hardware próximo ao alvo. (Putting the hardware in place. (Put it on very close to the TV sensor);)
 - Realize uma interação no teclado na TV. (On TV, go to a keyboard interaction. For example, a browser or another textbox interaction;)
 - Inicie o serIR. (It is time to start the serIR;)

### IR Tracking and Behavioral Profiling

### IR payload injection

### Sources and references

 - [NEC Infrared Codec Module (ver1.0) (YS-IRTM) - PDF] (https://aitendo3.sakura.ne.jp/aitendo_data/product_img/sensor/infrared/M1838-NEC-4P/M1838-NEC-4P_aitendo.pdf)
 - [MLT-BT05 4.0 Bluetooth module] (http://denethor.wlu.ca/arduino/MLT-BT05-AT-commands-TRANSLATED.pdf)
 - [WebOS simulator] (https://webostv.developer.lge.com/develop/tools/simulator-introduction)
 - [Smart TV LG 43" Full HD - HDR Ativo, webOS 4.5 ThinQ AI Processador Quad Core] (www.lg.com/br/tvs-e-soundbars/smart-tvs/43lm6300psb)

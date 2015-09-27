# Microbe
Experimental life form in Go language

Microbe is an experimental life form in Go language.

The objective is to create a small colony of autonomous robots on Linux compatible boards (Raspberry Pi, Olimex SOM, Beagle Board, etc) with :
- individual solar panel
- battery and/or supercapacitors
- way to monitor available energy and consumption
- motor to move the solar panel
- medium to comunicate between robots in the colony
- medium to share energy between robots

The objective of the colony is just to "stay alive" for now by :
- save energy of individual robots in case of low battery (by deactivating functions and/or services)
- move the individual solar panels to gain more energy
- transfer energy from a sacrified robot to another alive robots in case of crisis
- try to restart sacrified robots in the end of a crisis


The first exepriment will be done with only
- 3 robots
- no solar panels
- no motors
- no batery
- no energy monitoring
- no energy transfer
- ethernet cables

Fonctionality will be then added little by little from here.


- This is a part time project just for fun, so do not hope anything working before months
- We may, or not, also add later artificial inteligence stuff to the robots (energy saving and solar panel management)
- Robots may share their learnings and environment monitorings
- We may, or not, crypt the communication between robots
- We may, or not, add communication or interaction with the environment
- We may, or not, follow ROS.org architecture

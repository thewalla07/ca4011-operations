var X1 >= 0;
var X2 >= 0;
var X3 >= 0;
var X4 >= 0;
var X5 >= 0;
var X6 >= 0;

minimize Z:
    (4700*X1) + (500*X2) + (4700*X3) + (500*X4) + (4700*X5) + (500*X6);

subject to C1:
    X1 + X2 == 30;
subject to C2:
    (7*X1) - X3 - X4 == 20;
subject to C3:
    (7*X1) + (7*X3) -X5 - X6 == 70;
subject to C4:
    (7*X1) + (7*X3) + (7*X5) == 120;

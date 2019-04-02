(dynamic params/2).
(dynamic limitSup/1).
(dynamic limitInf/1).
(dynamic charSize/2).
(dynamic goal/3).
(dynamic platform/5).

%  debug input
debOutput(X, Y, Action, NextX, NextY) :-
    string_concat(X, ", ", CurPos),
    string_concat(CurPos, Y, CurPos2),
    write("Bird ("),
    write(CurPos2),
    write(") "),
    write(Action),
    write(" to "),
    string_concat(NextX, ", ", NextPos),
    string_concat(NextPos, NextY, NextPos2),
    write("("),
    write(NextPos2),
    write(")."),
    write("\n").

% err output
debErrOutput(X, Y, Action, Err) :-
    string_concat(X, ", ", CurPos),
    string_concat(CurPos, Y, CurPos2),
    write("Bird ("),
    write(CurPos2),
    write(") "),
    write(Action),
    write(" ("),
    write(Err),
    write(") "),
    write("\n").

yCalc(TSeq, Y) :-
    params(_, TFactor),
    TFactor >= 0.5,
    Y is (1.4283 * (TSeq^3)) - (7.7056 * (TSeq^2)) - (114.7708 * TSeq) + 122.0604, !.

yCalc(TSeq, Y) :-
    params(_, TFactor),
    TFactor >= 0.4,
    Y is (1.1619 * (TSeq^3)) - (5.834 * (TSeq^2)) - (93.1412 * TSeq) + 93.6452, !.

yCalc(TSeq, Y) :-
    params(_, TFactor),
    TFactor >= 0.3,
    Y is (0.8886 * (TSeq^3)) - (4.1712 * (TSeq^2)) - (70.7137 * TSeq) + 67.1291, !.

yCalc(TSeq, Y) :-
    params(_, TFactor),
    TFactor >= 0.2,
    Y is (0.4706 * (TSeq^3)) - (2.1024 * (TSeq^2)) - (48.2984 * TSeq) + 42.7892, !.

% bounds of the goal
goalBounds(MinX, MinY, MaxX, MaxY) :-
    goal(GX, GY, GR),
    MinX is GX-GR,
    MinY is GY-GR,
    MaxX is GX+GR,
    MaxY is GY+GR.

% bounds of the character
charBounds(X, Y, MinX, MinY, MaxX, MaxY) :-
    charSize(W, H),
    CW is W / 2,
    CH is H / 2,
    MinX is X-CW,
    MinY is Y-CH,
    MaxX is X+CW,
    MaxY is Y+CH.

% successful game
solve(X, Y, _, _, [none]) :-
    charBounds(X, Y, CMinX, CMinY, CMaxX, CMaxY),
    goalBounds(GMinX, GMinY, GMaxX, GMaxY),
    (   CMaxX>=GMinX,
        CMaxX=<GMaxX
    ;   CMinX=<GMaxX,
        CMinX>=GMinX
    ),
    (   CMaxY>=GMinY,
        CMaxY=<GMaxY
    ;   CMinY=<GMaxY,
        CMinY>=GMinY
    ),
    !, write("success\n").

% fail if it exceeds goal
solve(X, Y, _, Action, [Action]) :-
    params(RunSpeed, TFactor),
    NextX is X+(RunSpeed * TFactor / 2),
    goalBounds(_, _, MaxX, _),
    NextX>MaxX, !,
    debErrOutput(X, Y, Action, "exceeds goal").

% fail if it touches a top platform jumping
solve(X, Y, _, jump, [jump]) :-
    params(_, TFactor),
    TSeq is TFactor,
    yCalc(TSeq, YCalc),
    NextY is Y + YCalc,
    limitSup(LS),
    NextY>=LS,
    debErrOutput(X, Y, jump, "touches top base platform"), !.

% fail if it touches a base platform falling
solve(X, Y, TSeq, fall, [fall]) :-
    yCalc(TSeq, YCalc),
    NextY is Y + YCalc,
    limitInf(LI),
    NextY=<LI,
    debErrOutput(X, Y, fall, "touches bottom base platform"), !.

% fail if it touches a platform jumping
solve(X, Y, _, jump, [jump]) :-
    params(RunSpeed, TFactor),
    NextX is X+(RunSpeed * TFactor),
    TSeq is TFactor,
    yCalc(TSeq, YCalc),
    NextY is Y + YCalc,
    charBounds(NextX, NextY, CMinX, CMinY, CMaxX, CMaxY),
    (   platform(PMinX, PMinY, PMaxX, _, top),
        (   CMaxX>=PMinX,
            CMaxX<PMaxX
        ;   CMinX=<PMinX,
            CMaxX>=PMaxX
        ;   CMinX>=PMinX,
            CMinX=<PMaxX
        ),
        CMaxY>=PMinY
    ;   platform(PMinX, _, PMaxX, PMaxY, bottom),
        (   CMaxX>=PMinX,
            CMaxX<PMaxX
        ;   CMinX=<PMinX,
            CMaxX>=PMaxX
        ;   CMinX>=PMinX,
            CMinX=<PMaxX
        ),
        CMinY=<PMaxY
    ),
    debErrOutput(X, Y, jump, "touches platform jumping"), !.

% fail if it touches a platform falling
solve(X, Y, TSeq, fall, [fall]) :-
    params(RunSpeed, TFactor),
    NextX is X+(RunSpeed * TFactor),
    yCalc(TSeq, YCalc),
    NextY is Y + YCalc,
    charBounds(NextX, NextY, CMinX, CMinY, CMaxX, CMaxY),
    (   platform(PMinX, PMinY, PMaxX, _, top),
        (   CMaxX>=PMinX,
            CMaxX<PMaxX
        ;   CMinX=<PMinX,
            CMaxX>=PMaxX
        ;   CMinX>=PMinX,
            CMinX=<PMaxX
        ),
        CMaxY>=PMinY
    ;   platform(PMinX, _, PMaxX, PMaxY, bottom),
        (   CMaxX>=PMinX,
            CMaxX<PMaxX
        ;   CMinX=<PMinX,
            CMaxX>=PMaxX
        ;   CMinX>=PMinX,
            CMinX=<PMaxX
        ),
        CMinY=<PMaxY
    ), debErrOutput(X, Y, fall, "touches platform falling"), !.

% fall after fall
solve(X, Y, TSeq, fall, [fall|T]) :-
    params(RunSpeed, TFactor),
    NextX is X+(RunSpeed * TFactor),
    yCalc(TSeq, YCalc),
    NextY is Y + YCalc,
    writeln(Y),
    writeln(YCalc),
    NextTSeq is TSeq + TFactor,
    debOutput(X, Y, fall, NextX, NextY),
    solve(NextX, NextY, NextTSeq, fall, T).

% fall after jump
solve(X, Y, _, jump, [jump|T]) :-
    params(RunSpeed, TFactor),
    NextX is X+(RunSpeed * TFactor),
    TSeq is TFactor,
    yCalc(TSeq, YCalc),
    NextY is Y + YCalc,
    NextTSeq is TSeq + TFactor,
    debOutput(X, Y, jump, NextX, NextY),
    solve(NextX, NextY, NextTSeq, fall, T).

% jump after fall
solve(X, Y, TSeq, fall, [fall|T]) :-
    params(RunSpeed, TFactor),
    NextX is X+(RunSpeed * TFactor),
    yCalc(TSeq, YCalc),
    NextY is Y + YCalc,
    debOutput(X, Y, fall, NextX, NextY),
    solve(NextX, NextY, _, jump, T).

% jump after jump
solve(X, Y, _, jump, [jump|T]) :-
    params(RunSpeed, TFactor),
    NextX is X+(RunSpeed * TFactor),
    TSeq is TFactor,
    yCalc(TSeq, YCalc),
    NextY is Y + YCalc,
    debOutput(X, Y, jump, NextX, NextY),
    solve(NextX, NextY, _, jump, T).

% This is the code for all the exercises in the first document in prolog

% First Exercise

sumlist([], 0).
sumlist([X|Resto], S) :-
    sumlist(Resto, SResto), 
    S is X + SResto.

% Second Exercise


subconj([], _).

subconj([Elemento|RestoS1], S) :-
    member(Elemento, S), 
    subconj(RestoS1, S).  

% Third Exercise

aplanar([], []).

aplanar([Head|Tail], Flat) :-
    is_list(Head),
    aplanar(Head, FlatHead),
    aplanar(Tail, FlatTail),
    append(FlatHead, FlatTail, Flat).

aplanar([Head|Tail], [Head|FlatTail]) :-
    \+ is_list(Head),
    aplanar(Tail, FlatTail).


% Fourth Exercise

partir([], _, [], []).
partir([X|Resto], Umbral, [X|Menores], Mayores) :-
    X =< Umbral,
    partir(Resto, Umbral, Menores, Mayores).
partir([X|Resto], Umbral, Menores, [X|Mayores]) :-
    X > Umbral,
    partir(Resto, Umbral, Menores, Mayores).


% Fifth Exercise

contiene_subcadena(Cadena, Subcadena) :-
    sub_atom(Cadena, _, _, _, Subcadena).

filtrar_con_subcadena(ListaCadenas, Subcadena, CadenasFiltradas) :-
    include(contiene_subcadena(Subcadena), ListaCadenas, CadenasFiltradas).

% ------------------------- Third Class Exercise -----------------------------

% Define people with names, surnames, and chromosomes.
person(juan, perez, [1, 0, 1, 0, 1, 1, 0]).
person(maria, gomez, [0, 1, 1, 1, 0, 1, 0]).
person(luis, rodriguez, [1, 0, 1, 0, 1, 1, 1]).
person(ana, lopez, [0, 1, 1, 1, 1, 0, 0]).

% Define a predicate to calculate the compatibility between two people based on their chromosomes.
compatibility(Similarity, Person1, Person2) :-
    person(Person1, _, Chromosomes1),
    person(Person2, _, Chromosomes2),
    compare_chromosomes(Chromosomes1, Chromosomes2, 0, Similarity).

% Define a predicate to compare two lists of chromosomes and calculate the similarity.
compare_chromosomes([], [], Similarity, Similarity).
compare_chromosomes([X|Xs], [Y|Ys], Acc, Similarity) :-
    X =:= Y, % Compare the elements
    NewAcc is Acc + 1,
    compare_chromosomes(Xs, Ys, NewAcc, Similarity).
compare_chromosomes([_|Xs], [_|Ys], Acc, Similarity) :-
    compare_chromosomes(Xs, Ys, Acc, Similarity).

% Find compatible people for a given person.
find_compatible_people(Person, CompatibilityList) :-
    findall(Similarity-OtherPerson, (person(OtherPerson, _, _), Person \== OtherPerson, compatibility(Similarity, Person, OtherPerson)), CompatibilityList),
    keysort(CompatibilityList, SortedList),
    reverse(SortedList, ReversedList), % Sort in descending order
    maplist(=(Similarity-OtherPerson), ReversedList, CompatibilityList).


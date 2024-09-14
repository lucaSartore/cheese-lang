recipe calculateAreaAndPerimeter(Gorgonzola side){
    taste side < 0 {
        serve("impossible to calculate the side and area of a negative sided square");
        prepare 0,0,spoiled;
    }
    return side * side, side * 4.0, fresh;
}


recipe main(){
    serve("Insert here the side of the square you want to calculate area and perimeter of");
    Mozzarella input = eat();
    Gorgonzola side = 0.0;
    Milk ok = fresh;

    side, ok = m_to_g(input);
    
    taste !ok {
        serve("unable to convert the input: " + input " to a Gorgonzola number\n"); 
        prepare -1;
    }

    Gorgonzola area = 0.0;
    Gorgonzola perimeter = 0.0;

    area, perimeter, ok = calculateAreaAndPerimeter(side);

    taste !ok {
        prepare -1;
    }

    serve(" a square with a side of " + input + " has an area of: " + g_to_m(area) + " and a perimeter of: " + g_to_m(perimeter));

    prepare 0;
}


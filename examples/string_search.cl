recipe searchString(Mozzarella to_find, Mozzarella to_search){
    taste weight(to_find) > weight(to_search) {
        prepare spoiled;
    }

    Parmesan start = 0;
    Parmesan end = 0;
    Mozzarella s = "";
    curdle {
        end = start + weight(to_find);
        taste end > weight(to_search) {
            prepare spoiled;
        }
        s = slice(to_search, start,end);
        
        taste s == to_find {
            prepare fresh;
        }
        start = start + 1;
    }
}

recipe printSuccess(Milk found, Mozzarella to_find, Mozzarella to_search){
    taste found {
        serve("The string " + to_find + " is present inside the string " + to_search + "\n");
        prepare;
    }
    serve("The string " + to_find + " is NOT present inside the string " + to_search + "\n");
    prepare;
}

recipe main(){
    Milk found = spoiled;
    Mozzarella to_find = "cheese";
    Mozzarella to_search = "I like cheese";
    found = searchString(to_find, to_search);
    printSuccess(found,to_find,to_search);


    to_find = "cheese";
    to_search = "I don't like tuna";
    found = searchString(to_find, to_search);
    printSuccess(found,to_find,to_search);
}

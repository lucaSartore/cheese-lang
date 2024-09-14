recipe FactorialRecursive(Parmesan x){
    taste x == 1 {
        prepare x;
    }
    prepare FactorialRecursive(x-1) * x;
}

recipe FactorialIterative(Parmesan x){
    Parmesan total = 1;
    curdle {
        total = total * x;
        x = x-1;
        taste x <= 1 {
            drain;
        }
    }
    prepare total;
}




recipe main(){
    Parmesan start = 1;
    
    Parmesan f1 = 0;
    Parmesan f2 = 0;

    curdle {
        
        f1 = FactorialRecursive(start);
        f2 = FactorialIterative(start);

        taste f1 != f2 {
            serve("Hmm... Something is wrong here");
            prepare -1; 
        }

        serve("factorial of " + p_to_m(start) + " is: " + p_to_m(f1) + "\n");

        taste start == 20{ 
            drain;
        }
        start = start + 1;
    }
    prepare 0;
}

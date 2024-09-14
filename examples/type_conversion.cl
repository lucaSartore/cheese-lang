recipe main(){
    Milk ok = spoiled; 
    Parmesan a = 12;
    Gorgonzola b = 34.56;
    // "34.56"
    Mozzarella c = g_to_m(b);
    
    // 34
    Parmesan b2 = g_to_p(b);
    
    // "12"
    Mozzarella a2 = p_to_m(a);
    // "34"
    Mozzarella b3 = p_to_m(b2);

    // "12.34"
    Mozzarella d = a2 + "." + b3;
    
    // "34.5678"
    Mozzarella c2 = c + "78";

    // 34.5678
    Gorgonzola c3 = 0.0;
    c3, ok = m_to_g(c2);
    taste ok == spoiled {
        prepare -1;
    }
    
    Gorgonzola e = 0.0;

    e, ok = m_to_g("34.a");
    taste ok == fresh {
        prepare -1;
    }


    e, ok = m_to_g("34.0");
    taste ok == spoiled {
        prepare -1;
    }


    // 0.005678
    c3 = (c3 - e)/100.0;

    // 12.34
    Gorgonzola d2 = 0.0;
    d2, ok = m_to_g(d);

    taste ok == spoiled {
        prepare -1;
    }

    // 12.345678
    Gorgonzola final = d2 + c3;
    
    Mozzarella result = "Test 12.345678 == " + g_to_m(final) + "\n";
    
    serve(result);

    prepare 0;
}

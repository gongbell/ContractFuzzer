contract ClassicCheck {
    
    bool public classic;
 
    function ClassicCheck(){
        if (address(0xbf4ed7b27f1d666546e30d74d50d173d20bca754).balance > 1000000 ether)
            classic = false;
        else
            classic = true;
    }   
    
    function isClassic() constant returns (bool isClassic) {
        return classic;
    }
}
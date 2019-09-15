contract ClassicCheck {
       function isClassic() constant returns (bool isClassic);
}

contract SafeConditionalHFTransfer {

    bool classic;
    
    function SafeConditionalHFTransfer() {
        classic = ClassicCheck(0x882fb4240f9a11e197923d0507de9a983ed69239).isClassic();
    }
    
    function classicTransfer(address to) {
        if (!classic) 
            msg.sender.send(msg.value);
        else
            to.send(msg.value);
    }
    
    function transfer(address to) {
        if (classic)
            msg.sender.send(msg.value);
        else
            to.send(msg.value);
    }
    
}
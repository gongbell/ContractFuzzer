contract Agent{
    uint public count = 0;
    address  public call_contract_addr;
    bytes  public  call_msg_data;
    bool public turnoff = true;
    bool public hasValue = false;
    uint public sendCount = 0;
    uint public sendFailedCount =0;
    function() payable{
     if (turnoff){
        count ++;
        //turnoff set to false before statement "call_contract_addr.call.." rather than afer.
        //As we aim to test reentrancy only one times and  
        //more times for reentrancy is unnecessary.
        turnoff = false;
        call_contract_addr.call(call_msg_data);
        
     }else{
        turnoff = true;
     }
    }
    function Agent(){

    }
    function getContractAddr() returns(address addr){
        return call_contract_addr;
    }
    function getCallMsgData() returns(bytes msg_data){
        return call_msg_data;
    }
    function AgentCallWithoutValue(address contract_addr,bytes msg_data){
        hasValue = false;
        call_contract_addr  = contract_addr;
        call_msg_data = msg_data;
        contract_addr.call(msg_data);
    }
    function AgentCallWithValue(address contract_addr,bytes msg_data) payable{
      hasValue = true;
      uint msg_value = msg.value;
      call_contract_addr  = contract_addr;
      call_msg_data = msg_data;
      contract_addr.call.value(msg_value)(msg_data);
    }
    function AgentSend(address contract_addr) payable{
        sendCount ++;
        if (!contract_addr.send(msg.value))
            sendFailedCount++;
    }
}
pragma solidity 0.4.4;

contract Reward{
    
    function reward(uint32[] rewardsDistribution, address[] winners) payable{
        
        if(rewardsDistribution.length == 0 || rewardsDistribution.length > 64){ // do not risk gas shortage on reward
			throw;
		}
		//ensure rewardsDistribution give always something and do not give more to a lower scoring player
		uint32 prev = 0;
		for(uint8 i = 0; i < rewardsDistribution.length; i++){
			if(rewardsDistribution[i] == 0 ||  (prev != 0 && rewardsDistribution[i] > prev)){
				throw;
			}
			prev = rewardsDistribution[i];
		}
		
        uint8 numWinners = uint8(rewardsDistribution.length);

		if(numWinners > uint8(winners.length)){
			numWinners = uint8(winners.length);
		}
		
        uint forJack = msg.value;
		uint64 total = 0;
		for(uint8 j=0; j<numWinners; j++){ // distribute all the winning even if there is not enought winners
			total += rewardsDistribution[j];
		}
		for(uint8 k=0; k<numWinners; k++){
			uint value = (msg.value * rewardsDistribution[k]) / total;
			if(winners[k].send(value)){ // skip winner if fail to send but still use next distribution index
				forJack = forJack - value;
			}
		}
		
		if(forJack > 0){
		    if(!msg.sender.send(forJack)){
		        throw;
		    } 
		}
		
    }
    
}
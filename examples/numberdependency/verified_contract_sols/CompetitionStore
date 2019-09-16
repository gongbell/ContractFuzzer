/**
 *Submitted for verification at Etherscan.io on 2017-01-02
*/

/* Copyright (C) Etherplay <contact@etherplay.io> - All Rights Reserved */
pragma solidity 0.4.4;

contract CompetitionStore {
	
/////////////////////////////////////////////////////////////////// DATA /////////////////////////////////////////////////////////////
	
	//player's submission store the info required to verify its accuracy
	struct Submission{
		uint32 score; 
		uint32 durationRoundedDown; // duration in second of the game session
		uint32 version; // version of the game used
		uint64 seed; //seed used
		uint64 submitBlockNumber; // blockNumber at which the submission is processed
		bytes32 proofHash;//sha256 of proof : to save gas, the proof is not saved directly in the contract. Instead its hash is saved. The actual proof will be saved on a server. The player could potentially save it too. 
	}
	
	//player start game parameter
	struct Start{
		uint8 competitionIndex; //competition index (0 or 1) there is only 2 current competition per game, one is active, the other one being the older one which might have pending verification
		uint32 version;  //version of the game that the player score is based on
		uint64 seed; // the seed used for the game session
		uint64 time; // start time , used to check if the player is not taking too long to submit its score
	}
	
	// the values representing each competition
	struct Competition{
		uint8 numPastBlocks;// number of past block allowed, 1 is the minimum since you can only get the hash of a past block. Allow player to start play instantunously
		uint8 houseDivider; // how much the house takes : 4 means house take 1/4 (25%)
		uint16 lag; // define how much extra time is allowed to submit a score (to accomodate block time and delays)
		uint32 verificationWaitTime;// wait time allowed for submission past competition's end time 
		uint32 numPlayers;//current number of player that submited a score
		uint32 version; //the version of the game used for that competition, a hash of the code is published in the log upon changing
		uint32 previousVersion; // previousVersion to allow smooth update upon version change
		uint64 versionChangeBlockNumber; 
		uint64 switchBlockNumber; // the blockNumber at which the competition started
		uint64 endTime;//The time at which the competition is set to finish. No start can happen after that and the competition cannot be aborted before that
		uint88 price;  // the price for that competition, do not change 
		uint128 jackpot; // the current jackpot for that competition, this jackpot is then shared among the developer (in the deposit account for  funding development) and the winners (see houseDivider))
		uint32[] rewardsDistribution; // the length of it define how many winners there is and the distribution of the reward is the value for each index divided by the total
		mapping (address => Submission) submissions;  //only one submission per player per competition
		address[] players; // contain the list of players that submited a score for that competition
	}
		
	struct Game{
		mapping (address => Start) starts; // only 1 start per player, further override the current
		Competition[2] competitions; // 2 competitions only to save gas, overrite each other upon going to next competition
		uint8 currentCompetitionIndex; //can only be 1 or 0 (switch operation : 1 - currentCompetitionIndex)
	}

	mapping (string => Game) games;
	
	address organiser; // admin having control of the reward 
	address depositAccount;	 // is the receiver of the house part of the jackpot (see houseDivider) Can only be changed by the depositAccount.

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////



///////////////////////////////////////////////////////// EVENTS /////////////////////////////////////////////////////////////

	//event logging the hash of the game code for a particular version
	event VersionChange(
		string indexed gameID,
		uint32 indexed version,
		bytes32 codeHash // the sha256 of the game code as used by the player
	);

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////




//////////////////////////////////////////////////////// PLAYERS ACTIONS /////////////////////////////////////////////////////////////
	
	/*
	The seed is computed from the block hash and the sender address
	While the seed can be predicted for few block away (see : numPastBlocks) this is has no much relevance since a game session have a bigger duration,
	Remember this is not gambling game, this is a skill game, seed is only a small part of the game outcome
	*/
	function computeSeed(uint64 blockNumber, address player) internal constant returns(uint64 seed){ 
		return uint64(sha3(block.blockhash(blockNumber),block.blockhash(blockNumber-1),block.blockhash(blockNumber-2),block.blockhash(blockNumber-3),block.blockhash(blockNumber-4),block.blockhash(blockNumber-5),player)); 
	}
	
	/*
		probe the current state of the competition so player can start playing right away (need to commit a tx too to ensure its play will be considered though)
	*/
	function getSeedAndState(string gameID, address player) constant returns(uint64 seed, uint64 blockNumber, uint8 competitionIndex, uint32 version, uint64 endTime, uint88 price, uint32 myBestScore, uint64 competitionBlockNumber, uint64 registeredSeed){
		var game = games[gameID];

		competitionIndex = game.currentCompetitionIndex;
		var competition = game.competitions[competitionIndex];

		blockNumber = uint64(block.number-1);
		seed = computeSeed(blockNumber, player);
		version = competition.version;
		endTime = competition.endTime;
		price = competition.price;
		competitionBlockNumber = competition.switchBlockNumber;
		
		if (competition.submissions[player].submitBlockNumber >= competition.switchBlockNumber){
			myBestScore = competition.submissions[player].score;
		}else{
			myBestScore = 0;
		}
		
		registeredSeed = game.starts[player].seed;
	}
	
	
		
	function start(string gameID, uint64 blockNumber,uint8 competitionIndex, uint32 version) payable {
		var game = games[gameID];
		var competition = game.competitions[competitionIndex];

		if(msg.value != competition.price){
			throw;
		}

		if(
			competition.endTime <= now || //block play when time is up 
			competitionIndex != game.currentCompetitionIndex || //start happen just after a switch // should not be possible since endTime already ensure that a new competition cannot start before the end of the first
			version != competition.version && (version != competition.previousVersion || block.number > competition.versionChangeBlockNumber) || //ensure version is same as current (or previous if versionChangeBlockNumber is recent)
			block.number >= competition.numPastBlocks && block.number - competition.numPastBlocks > blockNumber //ensure start is not too old   
			){
				//if ether was sent, send it back if possible, else throw
				if(msg.value != 0 && !msg.sender.send(msg.value)){
					throw;
				}
				return;
		}
		
		competition.jackpot += uint128(msg.value); //increase the jackpot
		
		//save the start params
		game.starts[msg.sender] = Start({
			seed: computeSeed(blockNumber,msg.sender)
			, time : uint64(now)
			, competitionIndex : competitionIndex
			, version : version
		}); 
	}
		
	function submit(string gameID, uint64 seed, uint32 score, uint32 durationRoundedDown, bytes32 proofHash){ 
		var game = games[gameID];

		var gameStart = game.starts[msg.sender];
			
		//seed should be same, else it means double start and this one executing is from the old one 
		if(gameStart.seed != seed){
			return;
		}
		
		var competition = game.competitions[gameStart.competitionIndex];
		
		// game should not take too long to be submited
		if(now - gameStart.time > durationRoundedDown + competition.lag){ 
			return;
		}

		if(now >= competition.endTime + competition.verificationWaitTime){
			return; //this ensure verifier to get all the score at that time (should never be there though as game should ensure a maximumTime < verificationWaitTime)
		}
		
		var submission = competition.submissions[msg.sender];
		if(submission.submitBlockNumber < competition.switchBlockNumber){
			if(competition.numPlayers >= 4294967295){ //unlikely but if that happen this is for now the best place to stop
				return;
			}
		}else if (score <= submission.score){
			return;
		}
		
		var players = competition.players;
		//if player did not submit score yet => add player to list
		if(submission.submitBlockNumber < competition.switchBlockNumber){
			var currentNumPlayer = competition.numPlayers;
			if(currentNumPlayer >= players.length){
				players.push(msg.sender);
			}else{
				players[currentNumPlayer] = msg.sender;
			}
			competition.numPlayers = currentNumPlayer + 1;
		}
		
		competition.submissions[msg.sender] = Submission({
			proofHash:proofHash,
			seed:gameStart.seed,
			score:score,
			durationRoundedDown:durationRoundedDown,
			submitBlockNumber:uint64(block.number),
			version:gameStart.version
		});
		
	}
	
	/*
		accept donation payment : this increase the jackpot of the currentCompetition of the specified game
	*/
	function increaseJackpot(string gameID) payable{
		var game = games[gameID];
		game.competitions[game.currentCompetitionIndex].jackpot += uint128(msg.value); //extra ether is lost but this is not going to happen :)
	}

//////////////////////////////////////////////////////////////////////////////////////////

	
/////////////////////////////////////// PRIVATE ///////////////////////////////////////////
		
	function CompetitionStore(){
		organiser = msg.sender;
		depositAccount = msg.sender;
	}

	
	//give a starting jackpot by sending ether to the transaction
	function _startNextCompetition(string gameID, uint32 version, uint88 price, uint8 numPastBlocks, uint8 houseDivider, uint16 lag, uint64 duration, uint32 verificationWaitTime, bytes32 codeHash, uint32[] rewardsDistribution) payable{
		if(msg.sender != organiser){
			throw;
		}
		var game = games[gameID];
		var newCompetition = game.competitions[1 - game.currentCompetitionIndex]; 
		var currentCompetition = game.competitions[game.currentCompetitionIndex];
		//do not allow to switch if endTime is not over
		if(currentCompetition.endTime >= now){
			throw;
		}

		//block switch if reward was not called (numPlayers > 0)
		if(newCompetition.numPlayers > 0){
			throw;
		}
		
		if(houseDivider == 0){ 
			throw;
		}
		
		if(numPastBlocks < 1){
			throw;
		}
		
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

		if(version != currentCompetition.version){
			VersionChange(gameID,version,codeHash); 
		}
		
		game.currentCompetitionIndex = 1 - game.currentCompetitionIndex;
		
		newCompetition.switchBlockNumber = uint64(block.number);
		newCompetition.previousVersion = 0;
		newCompetition.versionChangeBlockNumber = 0;
		newCompetition.version = version;
		newCompetition.price = price; 
		newCompetition.numPastBlocks = numPastBlocks;
		newCompetition.rewardsDistribution = rewardsDistribution;
		newCompetition.houseDivider = houseDivider;
		newCompetition.lag = lag;
		newCompetition.jackpot += uint128(msg.value); //extra ether is lost but this is not going to happen :)
		newCompetition.endTime = uint64(now) + duration;
		newCompetition.verificationWaitTime = verificationWaitTime;
	}
	
	
	
	function _setBugFixVersion(string gameID, uint32 version, bytes32 codeHash, uint32 numBlockAllowedForPastVersion){
		if(msg.sender != organiser){
			throw;
		}

		var game = games[gameID];
		var competition = game.competitions[game.currentCompetitionIndex];
		
		if(version <= competition.version){ // a bug fix should be a new version (greater than previous version)
			throw;
		}
		
		if(competition.endTime <= now){ // cannot bugFix a competition that already ended
			return;
		}
		
		competition.previousVersion = competition.version;
		competition.versionChangeBlockNumber = uint64(block.number + numBlockAllowedForPastVersion);
		competition.version = version;
		VersionChange(gameID,version,codeHash);
	}

	function _setLagParams(string gameID, uint16 lag, uint8 numPastBlocks){
		if(msg.sender != organiser){
			throw;
		}
		
		if(numPastBlocks < 1){
			throw;
		}

		var game = games[gameID];
		var competition = game.competitions[game.currentCompetitionIndex];
		competition.numPastBlocks = numPastBlocks;
		competition.lag = lag;
	}

	function _rewardWinners(string gameID, uint8 competitionIndex, address[] winners){
		if(msg.sender != organiser){
			throw;
		}
		
		var competition = games[gameID].competitions[competitionIndex];

		//ensure time has passed so that players who started near the end can finish their session 
		//game should be made to ensure termination before verificationWaitTime, it is the game responsability
		if(int(now) - competition.endTime < competition.verificationWaitTime){
			throw;
		}

		
		if( competition.jackpot > 0){ // if there is no jackpot skip

			
			var rewardsDistribution = competition.rewardsDistribution;

			uint8 numWinners = uint8(rewardsDistribution.length);

			if(numWinners > uint8(winners.length)){
				numWinners = uint8(winners.length);
			}

			uint128 forHouse = competition.jackpot;
			if(numWinners > 0 && competition.houseDivider > 1){ //in case there is no winners (no players or only cheaters), the house takes all
				forHouse = forHouse / competition.houseDivider;
				uint128 forWinners = competition.jackpot - forHouse;

				uint64 total = 0;
				for(uint8 i=0; i<numWinners; i++){ // distribute all the winning even if there is not all the winners
					total += rewardsDistribution[i];
				}
				for(uint8 j=0; j<numWinners; j++){
					uint128 value = (forWinners * rewardsDistribution[j]) / total;
					if(!winners[j].send(value)){ // if fail give to house
						forHouse = forHouse + value;
					}
				}
			}
			
			if(!depositAccount.send(forHouse)){
				//in case sending to house failed 
				var nextCompetition = games[gameID].competitions[1 - competitionIndex];
				nextCompetition.jackpot = nextCompetition.jackpot + forHouse;	
			}

			
			competition.jackpot = 0;
		}
		
		
		competition.numPlayers = 0;
	}

	
	/*
		allow to change the depositAccount of the house share, only the depositAccount can change it, depositAccount == organizer at creation
	*/
	function _setDepositAccount(address newDepositAccount){
		if(depositAccount != msg.sender){
			throw;
		}
		depositAccount = newDepositAccount;
	}
	
	/*
		allow to change the organiser, in case this need be 
	*/
	function _setOrganiser(address newOrganiser){
		if(organiser != msg.sender){
			throw;
		}
		organiser = newOrganiser;
	}
	
	
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////// OTHER CONSTANT CALLS TO PROBE VALUES ////////////////////////////////////////////////////

	function getPlayerSubmissionFromCompetition(string gameID, uint8 competitionIndex, address playerAddress) constant returns(uint32 score, uint64 seed, uint32 duration, bytes32 proofHash, uint32 version, uint64 submitBlockNumber){
		var submission = games[gameID].competitions[competitionIndex].submissions[playerAddress];
		score = submission.score;
		seed = submission.seed;		
		duration = submission.durationRoundedDown;
		proofHash = submission.proofHash;
		version = submission.version;
		submitBlockNumber =submission.submitBlockNumber;
	}
	
	function getPlayersFromCompetition(string gameID, uint8 competitionIndex) constant returns(address[] playerAddresses, uint32 num){
		var competition = games[gameID].competitions[competitionIndex];
		playerAddresses = competition.players;
		num = competition.numPlayers;
	}

	function getCompetitionValues(string gameID, uint8 competitionIndex) constant returns (
		uint128 jackpot,
		uint88 price,
		uint32 version,
		uint8 numPastBlocks,
		uint64 switchBlockNumber,
		uint32 numPlayers,
		uint32[] rewardsDistribution,
		uint8 houseDivider,
		uint16 lag,
		uint64 endTime,
		uint32 verificationWaitTime,
		uint8 _competitionIndex
	){
		var competition = games[gameID].competitions[competitionIndex];
		jackpot = competition.jackpot;
		price = competition.price;
		version = competition.version;
		numPastBlocks = competition.numPastBlocks;
		switchBlockNumber = competition.switchBlockNumber;
		numPlayers = competition.numPlayers;
		rewardsDistribution = competition.rewardsDistribution;
		houseDivider = competition.houseDivider;
		lag = competition.lag;
		endTime = competition.endTime;
		verificationWaitTime = competition.verificationWaitTime;
		_competitionIndex = competitionIndex;
	}
	
	function getCurrentCompetitionValues(string gameID) constant returns (
		uint128 jackpot,
		uint88 price,
		uint32 version,
		uint8 numPastBlocks,
		uint64 switchBlockNumber,
		uint32 numPlayers,
		uint32[] rewardsDistribution,
		uint8 houseDivider,
		uint16 lag,
		uint64 endTime,
		uint32 verificationWaitTime,
		uint8 _competitionIndex
	)
	{
		return getCompetitionValues(gameID,games[gameID].currentCompetitionIndex);
	}
}

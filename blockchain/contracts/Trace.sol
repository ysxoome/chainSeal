pragma solidity ^0.4.25;

import "./Seal.sol";

contract Trace {
    struct SealData {
        Seal traceSeal;
        bool valid;
    }
    string _category;
    mapping(uint64 => SealData) private _seals;
    constructor(string  sealTp) public {
        _category = sealTp;
    }
    
    event newSealEvent(uint64 sealId);
    
    function createSeal(uint64 sealId) public returns(Seal) {
        require(!_seals[sealId].valid, "id really exists");
        
        _seals[sealId].valid = true;
        Seal traceSeal = new Seal(sealId);
        emit newSealEvent(sealId);
        _seals[sealId].traceSeal = traceSeal;
        return traceSeal;
    }
    
    function changeSealStatus(uint64 sealId, address sender, int16 SealStatus, string memory remark) public {
        require(_seals[sealId].valid, "not exists");
         _seals[sealId].traceSeal.changeStatus(sender, SealStatus, remark);
    }
      
     function getStatus(uint64 sealId) public view returns(int16) {
         require(_seals[sealId].valid, "not exists");
         return _seals[sealId].traceSeal.getStatus();
    }

     function getSeal(uint64 sealId) public view returns(Seal) {
         require(_seals[sealId].valid, "not exists");
         return _seals[sealId].traceSeal;
    }
}

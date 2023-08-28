pragma solidity ^0.4.25;

import "./Trace.sol";

contract TraceFactory {
    struct SealTrace {
        Trace  trace;
        bool valid;
    }
    mapping(string => SealTrace) private _sealCategory;

	event newTraceEvent(string sealGroup);
	
	//Create Trace commodity category
    function createTrace(string sealGroup) public returns(Trace) {
        require(!_sealCategory[sealGroup].valid, "The trademark already exists");
        Trace category = new Trace(sealGroup);
        _sealCategory[sealGroup].valid = true;
        _sealCategory[sealGroup].trace = category;
        emit newTraceEvent(sealGroup);
        return category;
    }
    
     function getTrace(string sealGroup) private view returns(Trace) {
        require(_sealCategory[sealGroup].valid, "The trademark has not exists");
        return _sealCategory[sealGroup].trace;
    }
    //Create Trace products    
    function createTraceSeal(string sealGroup, uint64 sealId) public returns(Seal) {
        Trace category = getTrace(sealGroup);
         return category.createSeal(sealId);
    }
    
    //Change product status
    function changeTraceSeal(string sealGroup, uint64 sealId, int16 sealStatus, string memory remark) public {
         Trace category = getTrace(sealGroup);
         category.changeSealStatus(sealId, msg.sender, sealStatus, remark);
    }
    
    //Query the current status of seal
     function getStatus(string sealGroup, uint64 sealId) public view returns(int16) {
         Trace category = getTrace(sealGroup);
         return category.getStatus(sealId);
    }
    
    //The whole process of querying seal
     function getTraceInfo(string sealGroup, uint64 sealId) public view returns(Seal) {
         Trace category = getTrace(sealGroup);
         return category.getSeal(sealId);
    }
    

}
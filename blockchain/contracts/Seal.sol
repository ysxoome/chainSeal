pragma solidity ^0.4.25;
pragma experimental ABIEncoderV2;

contract Seal {
    // 创建智能合约的地址是由以太坊网络生成的，与控制台的用户地址信息是无关的
    struct TraceData {
        address addr;     //Operator address
        int16 status;     //seal status
        uint timestamp;   //Operator time
        string remark;    //Digested Data
    }
    uint64 _sealId; 
    int16 _status;   //current status
    TraceData[] _traceData;
    
    event newStatus( address addr, int16 status, uint timestamp, string remark);
    
    constructor(uint64 sealId) public {
        _sealId = sealId;
        _traceData.push(TraceData({addr:msg.sender, status:0, timestamp:now, remark:"create"}));
        emit newStatus(msg.sender, 0, now, "create");
    }
    
    function changeStatus(address sender, int16 sealStatus, string memory remark) public {
        _status = sealStatus;
        _traceData.push(TraceData({addr:sender, status:sealStatus, timestamp:now, remark:remark}));
        emit newStatus(sender, sealStatus, now, remark);
    }
      
    function getStatus() public view returns(int16) {
        return _status;
    }
    
    function getTraceInfo() public view returns(TraceData[] memory _data) {
        return _traceData;
    }
    
}
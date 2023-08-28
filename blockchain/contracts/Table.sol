pragma solidity ^0.4.24;

contract TableFactory {
    function openTable(string) public view returns (Table); //open table

    function createTable(string, string, string) public returns (int256); //create table
}

//select condition
contract Condition {
    function EQ(string, int256) public view;

    function EQ(string, string) public view;

    function NE(string, int256) public view;

    function NE(string, string) public view;

    function GT(string, int256) public view;

    function GE(string, int256) public view;

    function LT(string, int256) public view;

    function LE(string, int256) public view;

    function limit(int256) public view;

    function limit(int256, int256) public view;
}

//one record
contract Entry {
    function getInt(string) public view returns (int256);

    function getString(string) public view returns (string);

    function set(string, int256) public;

    function set(string, string) public;
}

//record sets
contract Entries {
    function get(int256) public view returns (Entry);

    function size() public view returns (int256);
}

//Table main contract
contract Table {
    function select(string, Condition) public view returns (Entries);

    function insert(string, Entry) public returns (int256);

    function update(string, Entry, Condition) public returns (int256);

    function remove(string, Condition) public returns (int256);

    function newEntry() public view returns (Entry);

    function newCondition() public view returns (Condition);
}

contract ImageTableFactory {
    function openTable(string) public view returns (ImageTable);

    function createTable(string, string, string) public view returns (int256);
}

contract ImageTable {
    function get(string) public view returns (bool, Entry);

    function set(string, Entry) public returns (int256);

    function newEntry() public view returns (Entry);
}

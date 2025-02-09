// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract lyraAiX {

    struct PredictionInfo {
        address user;
        uint32 signalId;
        uint64 userId;
        uint8 choice;
        bool hasInvolved;
    }

    struct CheckInInfo {
        address user;
        uint32 taskId;
        uint64 userId;
        uint256 timestamp;
    }
    // current chain task id
    uint32 private chainTaskId;

    PredictionInfo[] public predictionList;

    CheckInInfo[] public checkInList;

    // Stores the last check-in timestamp for each user
    mapping(address => mapping(uint256 => CheckInInfo)) private checkIn;

    // Stores user's belief for each signal
    mapping(address => mapping(uint256 => PredictionInfo)) private userPrediction;

    // Event emitted when a user checks in
    event CheckinEvent(address indexed user, CheckInInfo info);

    // Event emitted when a user evaluates a signal
    event SignalPredictionEvent(address indexed user, PredictionInfo info);

    constructor(){}

    // Calculate the current day number (days since Unix epoch)
    function getCurrentDay() public view returns (uint256) {
        return block.timestamp / 86400; // 86400 seconds = 1 day
    }

    // Function for users to check in
    function checkin(address user, uint32 taskId, uint64 userId) external {
        require(msg.sender == user, "Only the user can checkin");
        uint256 currentDay = getCurrentDay();
        require(checkIn[user][currentDay].timestamp == 0, "Already checked in today");

        CheckInInfo memory info = CheckInInfo(user, taskId, userId, block.timestamp);
        checkIn[user][currentDay] = info;
        checkInList.push(info);
        // emit CheckinEvent(user, info);
    }

    function checkInResult(address user, uint256 day) external view returns(bool){
        return checkIn[user][day].timestamp > 0;
    }

    function signalPredict(address user, uint32 signalId, uint64 userId, uint8 choice) external {
        require(msg.sender == user, "Only the user can predict");
        require(!userPrediction[user][signalId].hasInvolved, "Voter has already voted");

        PredictionInfo memory info = PredictionInfo(user, signalId, userId, choice, true);
        userPrediction[user][signalId] = info;
        predictionList.push(info);
        //emit SignalPredictionEvent(user, info);
    }

    function signalPredictionResult(address user, uint32 signalId) external view returns (bool, uint32, uint8) {
        if (!userPrediction[user][signalId].hasInvolved) {
            return (false, 0, 0);
        }
        return (userPrediction[user][signalId].hasInvolved,
            userPrediction[user][signalId].signalId, userPrediction[user][signalId].choice);
    }

    function getCheckIns(uint256 lastIndex, uint256 length) public view returns (CheckInInfo[] memory) {
        require(lastIndex < checkInList.length, "Invalid lastIndex");

        uint256 endIndex = lastIndex + length;
        if (endIndex > checkInList.length) {
            endIndex = checkInList.length;
        }

        uint256 size = endIndex - lastIndex;
        CheckInInfo[] memory result = new CheckInInfo[](size);

        for (uint256 i = 0; i < size; i++) {
            result[i] = checkInList[lastIndex + i];
        }

        return result;
    }

    function getPredictions(uint256 lastIndex, uint256 length) public view returns (PredictionInfo[] memory) {
        require(lastIndex < predictionList.length, "Invalid lastIndex");

        uint256 endIndex = lastIndex + length;
        if (endIndex > predictionList.length) {
            endIndex = predictionList.length;
        }

        uint256 size = endIndex - lastIndex;
        PredictionInfo[] memory result = new PredictionInfo[](size);

        for (uint256 i = 0; i < size; i++) {
            result[i] = predictionList[lastIndex + i];
        }

        return result;
    }
}

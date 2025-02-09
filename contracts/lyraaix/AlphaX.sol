// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract lyraAiX {

    struct PredictionInfo {
        uint32 signalId;
        uint64 userId;
        uint8 choice;
        bool hasInvolved;
    }

    struct CheckInInfo {
        uint32 taskId;
        uint64 userId;
        uint256 timestamp;
    }
    // current chain task id
    uint32 private chainTaskId;

    // Stores the last check-in timestamp for each user
    mapping(address => mapping(uint256 => CheckInInfo)) private checkIn;

    // Stores user's belief for each signal
    mapping(address => mapping(uint256 => PredictionInfo)) private userPrediction;

    // Event emitted when a user checks in
    event CheckinEvent(address indexed user, CheckInInfo info);

    // Event emitted when a user evaluates a signal
    event SignalPredictionEvent(address indexed user, PredictionInfo info);

    constructor(uint32 id){
        chainTaskId = id;
    }

    // Calculate the current day number (days since Unix epoch)
    function getCurrentDay() public view returns (uint256) {
        return block.timestamp / 86400; // 86400 seconds = 1 day
    }

    // Function for users to check in
    function checkin(address user, uint32 taskId, uint64 userId) external {
        require(msg.sender == user, "Only the user can checkin");
        require(taskId == chainTaskId, "Invalid taskId");
        uint256 currentDay = getCurrentDay();
        require(checkIn[user][currentDay].timestamp == 0, "Already checked in today");

        CheckInInfo memory info = CheckInInfo(taskId, userId, block.timestamp);
        checkIn[user][currentDay] = info;
        emit CheckinEvent(user, info);
    }

    function checkInResult(address user, uint256 day) external view returns(bool){
        return checkIn[user][day].timestamp > 0;
    }

    function signalPredict(address user, uint32 signalId, uint64 userId, uint8 choice) external {
        require(msg.sender == user, "Only the user can predict");
        require(!userPrediction[user][signalId].hasInvolved, "Voter has already voted");

        PredictionInfo memory info = PredictionInfo(signalId, userId, choice, true);
        userPrediction[user][signalId] = info;

        emit SignalPredictionEvent(user, info);
    }

    function signalPredictionResult(address user, uint32 signalId) external view returns (bool, uint32, uint8) {
        if (!userPrediction[user][signalId].hasInvolved) {
            return (false, 0, 0);
        }
        return (userPrediction[user][signalId].hasInvolved,
            userPrediction[user][signalId].signalId, userPrediction[user][signalId].choice);
    }
}

namespace py user
namespace go user

enum Activity {
    ONLINE = 1,
    OFFLINE = 2
}

struct ActivityEvent {
    1: Activity activity,
    2: i64 timestamp,
    3: i64 userid
}

service UserActivity {
    void online(1:i64 userid, 2:ActivityEvent ev),

    void offline(1:i64 userid, 2:ActivityEvent ev)
}

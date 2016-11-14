

package awsapi


import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
)


func GetSession(profile string) (*session.Session, error) {
    var sess *session.Session
    var err error

    if profile != "" {
        sess, err = AuthAsProfile(profile)
    } else {
        sess = AuthAsInstance()
        err = nil
    }

    return sess, err
}


func AuthAsInstance() *session.Session {
    return session.New(aws.NewConfig())
}


func AuthAsProfile(profile string) (*session.Session, error) {
    sess, err := session.NewSessionWithOptions(
        session.Options{
            Profile: profile,
            SharedConfigState: session.SharedConfigEnable,
        })
    return sess, err
}

// func withCredentials(id string, key string, region string) (*session.Session, error) {
//     sess, err := session.NewSession(
//         &aws.Config{
//             Region: &region,
//             Credentials: credentials.NewStaticCredentials(id, key, ""),
//         })
//     return sess, err
// }

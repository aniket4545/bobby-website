//this file contains all the call related to AWS and common functionality
import AWS from 'aws-sdk'; 

const aws_bucket = '';
const aws_region = '';
const aws_identityPoolId = '';

function uploadToAWS(file){
    return new Promise((resolve, reject)=>{
        //give a call to backend which will return some radomid/libid

        AWS.config.update({
            region: aws_region,
            credentials: new AWS.CognitoIdentityCredentials({
                IdentityPoolId: aws_identityPoolId
            })
        });
        // bucket would be some bucket with library id bucket/libraryid/contentid.ext
        var s3 = new AWS.S3({params:{Bucket: aws_bucket + ''}});
        var params = { 
            Key: '',
            Body: file,
            ACL: 'public-read',
            ContentType: ''
        };
        s3.upload({params}, function(error,data){
            if (err){
                reject(error);
            }else{
                //give call to backend to save image inside image map
                //against some library --> libid to save image url against the panel
                resolve(data);
            }
        });
    })       
}
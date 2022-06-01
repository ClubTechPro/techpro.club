package mailers

import (
	"fmt"
	"net/http"
	"sources/common"

	//go get -u github.com/aws/aws-sdk-go
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
    // Replace sender@example.com with your "From" address. 
    // This address must be verified with Amazon SES.
    Sender = "Chilarai from Techpro.Club<hello@techpro.club>"
    
    // Replace recipient@example.com with a "To" address. If your account 
    // is still in the sandbox, this address must be verified.
    Recipient = "hello@techpro.club"

    // Specify a configuration set. To use a configuration
    // set, comment the next line and line 92.
    //ConfigurationSet = "ConfigSet"
    
    // The subject line for the email.
    Subject = "Welcome to Techpro.Club"
    
    // The HTML body for the email.
    HtmlBody =  "<p>Hello</p>" +
		"<p>I am Chilarai, co-founder of Techpro.club. Thank you for joining us. We are a community of opensource contributors, which helps you to discover wonderful open source projects and encourages you to participate in the holistic growth of the projects as well as your portfolio.</p>" +
		
		"<h3>Where to head next?</h3>" +
		"<p>We are currently in Beta stage and onboarding more projects and contributors. But in the meanwhile, you can</p>" +
		"<ul>" +
		"<li>Fill up the Contributor preferences to get notified when you have matching projects</li>"+ 
		"<li>If you have an opensource project and looking for contributors, head to <a href='https://techpro.club/projects'>https://techpro.club/projects</a> and post to our community</li>"+ 
		"<li>Hit a star on our repository or fork it to contribute on <a href='https://github.com/ClubTechPro/techpro.club'>Github</a></li>"+ 
		"<li>Follow us and give us a shout out on <a href='https://twitter.com/ClubTechpro'>Twitter @ClubTechpro</a></li>"+ 
		"<li>Read our blogs <a href='https://blogs.techpro.club'>https://blogs.techpro.club</a>.</li>"+ 
		"<li>Help our community grow by spreading the word with your friends and family.</li>"+ 
		"<li>Lastly watch out for our emails for more details</li>"+ 
		"</ul>" + 
		"<br/><br/>Feel free to write back<br/><br/>" + 
		"<p>Best</p>" +
		"<p><b>Chilarai,</b><br/>Co founder, <a href='https://techpro.club'>Techpro.club</a><br/>" + 
		"<a href='https://github.com/chilarai'>Github</a>, <a href='https://twitter.com/chilly5476'>Twitter</a>";

    
    //The email body for recipients with non-HTML email clients.
    TextBody = "This email was sent with Amazon SES using the AWS SDK for Go."
    
    // The character encoding for the email.
    CharSet = "UTF-8"
)

func RegistrationEmail(w http.ResponseWriter, r *http.Request) {
    // Create a new session in the us-west-2 region.
    // Replace us-west-2 with the AWS Region you're using for Amazon SES.

	sesRegion := common.GetSesRegion()
	sesAccessID := common.GetSesAccessID()
	sesSecretKey := common.GetSesSecretKey()

    sess, err := session.NewSession(&aws.Config{
        Region:aws.String(sesRegion), 
		Credentials:credentials.NewStaticCredentials(sesAccessID,sesSecretKey,""),},
		
    )

	if err != nil {
		fmt.Println(err.Error())
	} else {
		// Create an SES session.
		svc := ses.New(sess)
			
		// Assemble the email.
		input := &ses.SendEmailInput{
			Destination: &ses.Destination{
				CcAddresses: []*string{
				},
				ToAddresses: []*string{
					aws.String(Recipient),
				},
			},
			Message: &ses.Message{
				Body: &ses.Body{
					Html: &ses.Content{
						Charset: aws.String(CharSet),
						Data:    aws.String(HtmlBody),
					},
					Text: &ses.Content{
						Charset: aws.String(CharSet),
						Data:    aws.String(TextBody),
					},
				},
				Subject: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(Subject),
				},
			},
			Source: aws.String(Sender),
				// Uncomment to use a configuration set
				//ConfigurationSetName: aws.String(ConfigurationSet),
		}

		// Attempt to send the email.
		result, err := svc.SendEmail(input)

		// Display error messages if they occur.
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case ses.ErrCodeMessageRejected:
					fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
				case ses.ErrCodeMailFromDomainNotVerifiedException:
					fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
				case ses.ErrCodeConfigurationSetDoesNotExistException:
					fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}

			return
		}

		fmt.Println("Email Sent to address: " + Recipient)
    	fmt.Println(result)
	}
}
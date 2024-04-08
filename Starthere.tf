provider "aws" {
  region = var.aws_region
}

resource "aws_instance" "example" {
  ami           = var.ami
  instance_type = var.instance_type

  tags = {
    Name = var.instance_name
  }
}



//provider "aws" {
//  region = "us-west-2"  # Specify your desired AWS region
//}

//resource "aws_instance" "example" {
//  ami           = "ami-0c55b159cbfafe1f0"  # Specify the AMI ID of the desired OS
//  instance_type = "t2.micro"               # Specify the instance type
//
//  tags = {
//    Name = "ExampleInstance"                # Tag the instance with a name
//  }
//}

//This Terraform script does the following:

//Specifies the AWS provider and the desired region (us-west-2 in this example).
//Defines an aws_instance resource named "example".
//Specifies the AMI ID of the desired operating system (Ubuntu 20.04 LTS in this case) using the ami parameter.
//Sets the instance type to t2.micro.
//Tags the instance with a name "ExampleInstance".
//You'll need to have AWS credentials configured either through environment variables, shared credentials file,
//or instance profile with appropriate permissions to create EC2 instances. Additionally, ensure that Terraform is installed on your system
//and properly configured.
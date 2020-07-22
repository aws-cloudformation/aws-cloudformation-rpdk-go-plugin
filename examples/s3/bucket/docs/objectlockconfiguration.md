# Example::S3::Bucket ObjectLockConfiguration

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#objectlockenabled" title="ObjectLockEnabled">ObjectLockEnabled</a>" : <i>String</i>,
    "<a href="#rule" title="Rule">Rule</a>" : <i><a href="objectlockrule.md">ObjectLockRule</a></i>
}
</pre>

### YAML

<pre>
<a href="#objectlockenabled" title="ObjectLockEnabled">ObjectLockEnabled</a>: <i>String</i>
<a href="#rule" title="Rule">Rule</a>: <i><a href="objectlockrule.md">ObjectLockRule</a></i>
</pre>

## Properties

#### ObjectLockEnabled

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Rule

_Required_: No

_Type_: <a href="objectlockrule.md">ObjectLockRule</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)


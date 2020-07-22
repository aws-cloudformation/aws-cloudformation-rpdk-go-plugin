# Example::S3::Bucket ReplicationConfiguration

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#role" title="Role">Role</a>" : <i>String</i>,
    "<a href="#rules" title="Rules">Rules</a>" : <i>[ <a href="replicationrule.md">ReplicationRule</a>, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#role" title="Role">Role</a>: <i>String</i>
<a href="#rules" title="Rules">Rules</a>: <i>
      - <a href="replicationrule.md">ReplicationRule</a></i>
</pre>

## Properties

#### Role

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Rules

_Required_: Yes

_Type_: List of <a href="replicationrule.md">ReplicationRule</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)


#!/usr/bin/perl -w
#
# Write comments to /srv/comments on the local host.
#
# Steve



use strict;
use warnings;

use CGI;
use Encode 'decode_utf8';
use Text::Markdown;


#
#  The directory beneath which we store the comments on-disk.
#
my $COMMENT = "/srv/comments/";

#
#  Get the parameters from the request.
#
my $cgi = new CGI();

my $name = $cgi->param('name') || undef;
$name = decode_utf8($name) if ($name);

my $mail = $cgi->param('mail') || undef;
$mail = decode_utf8($mail) if ($mail);

my $body = $cgi->param('body') || undef;
$body = decode_utf8($body) if ($body);

my $id = $cgi->param('id') || undef;
$id = decode_utf8($id) if ($id);

my $link = $cgi->param('link') || undef;
$link = decode_utf8($link) if ($link);
$link = lc($link)          if ($link);

my $cap = $cgi->param('robot') || undef;


#
# Strip newlines
#
$link =~ s/[\r\n]//g if ($link);
$id =~ s/[\r\n]//g   if ($id);
$name =~ s/[\r\n]//g if ($name);
$mail =~ s/[\r\n]//g if ($mail);


#
#  If any are missing just redirect back to the blog homepage.
#
if ( !defined($name) ||
     !length($name)  ||
     !defined($mail) ||
     !length($mail)  ||
     !defined($body) ||
     !length($body)  ||
     !defined($id)   ||
     !length($id) )
{
    print "Location: http://" . $ENV{ 'HTTP_HOST' } . "/#missing-field\n\n";
    exit;
}


#
#  Does the captcha value contain text?  If so spam.
#
if ( defined($cap) && length($cap) )
{
    print "Location: http://" . $ENV{ 'HTTP_HOST' } . "/#robot\n\n";
    exit;
}


#
#  Convert the message to crude HTML.
#
$body =~ s/\n$/<br>\n/mg;

#
#  Otherwise save them away.
#
#
#  ID.
#
if ( $id =~ /^(.*)[\/\\](.*)$/ )
{
    $id = $2;
}


#
#  Show the header
#
print "Content-type: text/html\n\n";


#
# get the current time - in seconds-past the epoch
#
my $timestr = scalar( time() );


#
#  Open the file.
#
my $file = $COMMENT . "/" . $id . "." . $timestr;
$file =~ s/[ \t]//g;
open( FILE, ">:encoding(UTF-8)", $file );
print FILE "Name: $name\n";
print FILE "Mail: $mail\n";
print FILE "Link: $link\n" if ( defined($link) );
print FILE "User-Agent: $ENV{'HTTP_USER_AGENT'}\n";
print FILE "IP-Address: $ENV{'REMOTE_ADDR'}\n";
print FILE "\n";

#
# Convert to HTML, if we can.
#
my $html = Text::Markdown::markdown($body);

#
# Save the output to the comment-file.
#
print FILE $html;
close(FILE);

#
#  Now show the user the thanks message..
#
print <<EOF;
<html>
 <head>
  <title>Thanks For Your Comment</title>
 </head>
 <body>
  <h2>Thanks!</h2>
  <p>Your comment will be included the next time this blog is rebuilt.</p>
  <p><a href="https://blog.steve.fi/">Return to blog</a>.</p>
 </body>
</html>
EOF

exit 0;

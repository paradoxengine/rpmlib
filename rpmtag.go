package rpmlib

/*
#include <rpm/rpmtag.h>
*/
import "C"

type RpmTagVal C.rpmTag

const (
	RPMTAG_HEADERIMAGE       RpmTagVal = C.RPMTAG_HEADERIMAGE
	RPMTAG_HEADERSIGNATURES  RpmTagVal = C.RPMTAG_HEADERSIGNATURES
	RPMTAG_HEADERIMMUTABLE   RpmTagVal = C.RPMTAG_HEADERIMMUTABLE
	RPMTAG_HEADERREGIONS     RpmTagVal = C.RPMTAG_HEADERREGIONS
	RPMTAG_HEADERI18NTABLE   RpmTagVal = C.RPMTAG_HEADERI18NTABLE
	RPMTAG_SIG_BASE          RpmTagVal = C.RPMTAG_SIG_BASE
	RPMTAG_SIGSIZE           RpmTagVal = C.RPMTAG_SIGSIZE
	RPMTAG_SIGLEMD5_1        RpmTagVal = C.RPMTAG_SIGLEMD5_1
	RPMTAG_SIGPGP            RpmTagVal = C.RPMTAG_SIGPGP
	RPMTAG_SIGLEMD5_2        RpmTagVal = C.RPMTAG_SIGLEMD5_2
	RPMTAG_SIGMD5            RpmTagVal = C.RPMTAG_SIGMD5
	RPMTAG_PKGID             RpmTagVal = C.RPMTAG_PKGID
	RPMTAG_SIGGPG            RpmTagVal = C.RPMTAG_SIGGPG
	RPMTAG_SIGPGP5           RpmTagVal = C.RPMTAG_SIGPGP5
	RPMTAG_BADSHA1_1         RpmTagVal = C.RPMTAG_BADSHA1_1
	RPMTAG_BADSHA1_2         RpmTagVal = C.RPMTAG_BADSHA1_2
	RPMTAG_PUBKEYS           RpmTagVal = C.RPMTAG_PUBKEYS
	RPMTAG_DSAHEADER         RpmTagVal = C.RPMTAG_DSAHEADER
	RPMTAG_RSAHEADER         RpmTagVal = C.RPMTAG_RSAHEADER
	RPMTAG_SHA1HEADER        RpmTagVal = C.RPMTAG_SHA1HEADER
	RPMTAG_HDRID             RpmTagVal = C.RPMTAG_HDRID
	RPMTAG_LONGSIGSIZE       RpmTagVal = C.RPMTAG_LONGSIGSIZE
	RPMTAG_LONGARCHIVESIZE   RpmTagVal = C.RPMTAG_LONGARCHIVESIZE
	RPMTAG_NAME              RpmTagVal = C.RPMTAG_NAME
	RPMTAG_N                 RpmTagVal = C.RPMTAG_N
	RPMTAG_VERSION           RpmTagVal = C.RPMTAG_VERSION
	RPMTAG_V                 RpmTagVal = C.RPMTAG_V
	RPMTAG_RELEASE           RpmTagVal = C.RPMTAG_RELEASE
	RPMTAG_R                 RpmTagVal = C.RPMTAG_R
	RPMTAG_EPOCH             RpmTagVal = C.RPMTAG_EPOCH
	RPMTAG_E                 RpmTagVal = C.RPMTAG_E
	RPMTAG_SUMMARY           RpmTagVal = C.RPMTAG_SUMMARY
	RPMTAG_DESCRIPTION       RpmTagVal = C.RPMTAG_DESCRIPTION
	RPMTAG_BUILDTIME         RpmTagVal = C.RPMTAG_BUILDTIME
	RPMTAG_BUILDHOST         RpmTagVal = C.RPMTAG_BUILDHOST
	RPMTAG_INSTALLTIME       RpmTagVal = C.RPMTAG_INSTALLTIME
	RPMTAG_SIZE              RpmTagVal = C.RPMTAG_SIZE
	RPMTAG_DISTRIBUTION      RpmTagVal = C.RPMTAG_DISTRIBUTION
	RPMTAG_VENDOR            RpmTagVal = C.RPMTAG_VENDOR
	RPMTAG_GIF               RpmTagVal = C.RPMTAG_GIF
	RPMTAG_XPM               RpmTagVal = C.RPMTAG_XPM
	RPMTAG_LICENSE           RpmTagVal = C.RPMTAG_LICENSE
	RPMTAG_PACKAGER          RpmTagVal = C.RPMTAG_PACKAGER
	RPMTAG_GROUP             RpmTagVal = C.RPMTAG_GROUP
	RPMTAG_CHANGELOG         RpmTagVal = C.RPMTAG_CHANGELOG
	RPMTAG_SOURCE            RpmTagVal = C.RPMTAG_SOURCE
	RPMTAG_PATCH             RpmTagVal = C.RPMTAG_PATCH
	RPMTAG_URL               RpmTagVal = C.RPMTAG_URL
	RPMTAG_OS                RpmTagVal = C.RPMTAG_OS
	RPMTAG_ARCH              RpmTagVal = C.RPMTAG_ARCH
	RPMTAG_PREIN             RpmTagVal = C.RPMTAG_PREIN
	RPMTAG_POSTIN            RpmTagVal = C.RPMTAG_POSTIN
	RPMTAG_PREUN             RpmTagVal = C.RPMTAG_PREUN
	RPMTAG_POSTUN            RpmTagVal = C.RPMTAG_POSTUN
	RPMTAG_OLDFILENAMES      RpmTagVal = C.RPMTAG_OLDFILENAMES
	RPMTAG_FILESIZES         RpmTagVal = C.RPMTAG_FILESIZES
	RPMTAG_FILESTATES        RpmTagVal = C.RPMTAG_FILESTATES
	RPMTAG_FILEMODES         RpmTagVal = C.RPMTAG_FILEMODES
	RPMTAG_FILEUIDS          RpmTagVal = C.RPMTAG_FILEUIDS
	RPMTAG_FILEGIDS          RpmTagVal = C.RPMTAG_FILEGIDS
	RPMTAG_FILERDEVS         RpmTagVal = C.RPMTAG_FILERDEVS
	RPMTAG_FILEMTIMES        RpmTagVal = C.RPMTAG_FILEMTIMES
	RPMTAG_FILEDIGESTS       RpmTagVal = C.RPMTAG_FILEDIGESTS
	RPMTAG_FILEMD5S          RpmTagVal = C.RPMTAG_FILEMD5S
	RPMTAG_FILELINKTOS       RpmTagVal = C.RPMTAG_FILELINKTOS
	RPMTAG_FILEFLAGS         RpmTagVal = C.RPMTAG_FILEFLAGS
	RPMTAG_ROOT              RpmTagVal = C.RPMTAG_ROOT
	RPMTAG_FILEUSERNAME      RpmTagVal = C.RPMTAG_FILEUSERNAME
	RPMTAG_FILEGROUPNAME     RpmTagVal = C.RPMTAG_FILEGROUPNAME
	RPMTAG_EXCLUDE           RpmTagVal = C.RPMTAG_EXCLUDE
	RPMTAG_EXCLUSIVE         RpmTagVal = C.RPMTAG_EXCLUSIVE
	RPMTAG_ICON              RpmTagVal = C.RPMTAG_ICON
	RPMTAG_SOURCERPM         RpmTagVal = C.RPMTAG_SOURCERPM
	RPMTAG_FILEVERIFYFLAGS   RpmTagVal = C.RPMTAG_FILEVERIFYFLAGS
	RPMTAG_ARCHIVESIZE       RpmTagVal = C.RPMTAG_ARCHIVESIZE
	RPMTAG_PROVIDENAME       RpmTagVal = C.RPMTAG_PROVIDENAME
	RPMTAG_PROVIDES          RpmTagVal = C.RPMTAG_PROVIDES
	RPMTAG_P                 RpmTagVal = C.RPMTAG_P
	RPMTAG_REQUIREFLAGS      RpmTagVal = C.RPMTAG_REQUIREFLAGS
	RPMTAG_REQUIRENAME       RpmTagVal = C.RPMTAG_REQUIRENAME
	RPMTAG_REQUIRES          RpmTagVal = C.RPMTAG_REQUIRES
	RPMTAG_REQUIREVERSION    RpmTagVal = C.RPMTAG_REQUIREVERSION
	RPMTAG_NOSOURCE          RpmTagVal = C.RPMTAG_NOSOURCE
	RPMTAG_NOPATCH           RpmTagVal = C.RPMTAG_NOPATCH
	RPMTAG_CONFLICTFLAGS     RpmTagVal = C.RPMTAG_CONFLICTFLAGS
	RPMTAG_CONFLICTNAME      RpmTagVal = C.RPMTAG_CONFLICTNAME
	RPMTAG_CONFLICTS         RpmTagVal = C.RPMTAG_CONFLICTS
	RPMTAG_C                 RpmTagVal = C.RPMTAG_C
	RPMTAG_CONFLICTVERSION   RpmTagVal = C.RPMTAG_CONFLICTVERSION
	RPMTAG_DEFAULTPREFIX     RpmTagVal = C.RPMTAG_DEFAULTPREFIX
	RPMTAG_BUILDROOT         RpmTagVal = C.RPMTAG_BUILDROOT
	RPMTAG_INSTALLPREFIX     RpmTagVal = C.RPMTAG_INSTALLPREFIX
	RPMTAG_EXCLUDEARCH       RpmTagVal = C.RPMTAG_EXCLUDEARCH
	RPMTAG_EXCLUDEOS         RpmTagVal = C.RPMTAG_EXCLUDEOS
	RPMTAG_EXCLUSIVEARCH     RpmTagVal = C.RPMTAG_EXCLUSIVEARCH
	RPMTAG_EXCLUSIVEOS       RpmTagVal = C.RPMTAG_EXCLUSIVEOS
	RPMTAG_AUTOREQPROV       RpmTagVal = C.RPMTAG_AUTOREQPROV
	RPMTAG_RPMVERSION        RpmTagVal = C.RPMTAG_RPMVERSION
	RPMTAG_TRIGGERSCRIPTS    RpmTagVal = C.RPMTAG_TRIGGERSCRIPTS
	RPMTAG_TRIGGERNAME       RpmTagVal = C.RPMTAG_TRIGGERNAME
	RPMTAG_TRIGGERVERSION    RpmTagVal = C.RPMTAG_TRIGGERVERSION
	RPMTAG_TRIGGERFLAGS      RpmTagVal = C.RPMTAG_TRIGGERFLAGS
	RPMTAG_TRIGGERINDEX      RpmTagVal = C.RPMTAG_TRIGGERINDEX
	RPMTAG_VERIFYSCRIPT      RpmTagVal = C.RPMTAG_VERIFYSCRIPT
	RPMTAG_CHANGELOGTIME     RpmTagVal = C.RPMTAG_CHANGELOGTIME
	RPMTAG_CHANGELOGNAME     RpmTagVal = C.RPMTAG_CHANGELOGNAME
	RPMTAG_CHANGELOGTEXT     RpmTagVal = C.RPMTAG_CHANGELOGTEXT
	RPMTAG_BROKENMD5         RpmTagVal = C.RPMTAG_BROKENMD5
	RPMTAG_PREREQ            RpmTagVal = C.RPMTAG_PREREQ
	RPMTAG_PREINPROG         RpmTagVal = C.RPMTAG_PREINPROG
	RPMTAG_POSTINPROG        RpmTagVal = C.RPMTAG_POSTINPROG
	RPMTAG_PREUNPROG         RpmTagVal = C.RPMTAG_PREUNPROG
	RPMTAG_POSTUNPROG        RpmTagVal = C.RPMTAG_POSTUNPROG
	RPMTAG_BUILDARCHS        RpmTagVal = C.RPMTAG_BUILDARCHS
	RPMTAG_OBSOLETENAME      RpmTagVal = C.RPMTAG_OBSOLETENAME
	RPMTAG_OBSOLETES         RpmTagVal = C.RPMTAG_OBSOLETES
	RPMTAG_O                 RpmTagVal = C.RPMTAG_O
	RPMTAG_VERIFYSCRIPTPROG  RpmTagVal = C.RPMTAG_VERIFYSCRIPTPROG
	RPMTAG_TRIGGERSCRIPTPROG RpmTagVal = C.RPMTAG_TRIGGERSCRIPTPROG
	RPMTAG_DOCDIR            RpmTagVal = C.RPMTAG_DOCDIR
	RPMTAG_COOKIE            RpmTagVal = C.RPMTAG_COOKIE
	RPMTAG_FILEDEVICES       RpmTagVal = C.RPMTAG_FILEDEVICES
	RPMTAG_FILEINODES        RpmTagVal = C.RPMTAG_FILEINODES
	RPMTAG_FILELANGS         RpmTagVal = C.RPMTAG_FILELANGS
	RPMTAG_PREFIXES          RpmTagVal = C.RPMTAG_PREFIXES
	RPMTAG_INSTPREFIXES      RpmTagVal = C.RPMTAG_INSTPREFIXES
	RPMTAG_TRIGGERIN         RpmTagVal = C.RPMTAG_TRIGGERIN
	RPMTAG_TRIGGERUN         RpmTagVal = C.RPMTAG_TRIGGERUN
	RPMTAG_TRIGGERPOSTUN     RpmTagVal = C.RPMTAG_TRIGGERPOSTUN
	RPMTAG_AUTOREQ           RpmTagVal = C.RPMTAG_AUTOREQ
	RPMTAG_AUTOPROV          RpmTagVal = C.RPMTAG_AUTOPROV
	RPMTAG_CAPABILITY        RpmTagVal = C.RPMTAG_CAPABILITY
	RPMTAG_SOURCEPACKAGE     RpmTagVal = C.RPMTAG_SOURCEPACKAGE
	RPMTAG_OLDORIGFILENAMES  RpmTagVal = C.RPMTAG_OLDORIGFILENAMES
	RPMTAG_BUILDPREREQ       RpmTagVal = C.RPMTAG_BUILDPREREQ
	RPMTAG_BUILDREQUIRES     RpmTagVal = C.RPMTAG_BUILDREQUIRES
	RPMTAG_BUILDCONFLICTS    RpmTagVal = C.RPMTAG_BUILDCONFLICTS
	RPMTAG_BUILDMACROS       RpmTagVal = C.RPMTAG_BUILDMACROS
	RPMTAG_PROVIDEFLAGS      RpmTagVal = C.RPMTAG_PROVIDEFLAGS
	RPMTAG_PROVIDEVERSION    RpmTagVal = C.RPMTAG_PROVIDEVERSION
	RPMTAG_OBSOLETEFLAGS     RpmTagVal = C.RPMTAG_OBSOLETEFLAGS
	RPMTAG_OBSOLETEVERSION   RpmTagVal = C.RPMTAG_OBSOLETEVERSION
	RPMTAG_DIRINDEXES        RpmTagVal = C.RPMTAG_DIRINDEXES
	RPMTAG_BASENAMES         RpmTagVal = C.RPMTAG_BASENAMES
	RPMTAG_DIRNAMES          RpmTagVal = C.RPMTAG_DIRNAMES
	RPMTAG_ORIGDIRINDEXES    RpmTagVal = C.RPMTAG_ORIGDIRINDEXES
	RPMTAG_ORIGBASENAMES     RpmTagVal = C.RPMTAG_ORIGBASENAMES
	RPMTAG_ORIGDIRNAMES      RpmTagVal = C.RPMTAG_ORIGDIRNAMES
	RPMTAG_OPTFLAGS          RpmTagVal = C.RPMTAG_OPTFLAGS
	RPMTAG_DISTURL           RpmTagVal = C.RPMTAG_DISTURL
	RPMTAG_PAYLOADFORMAT     RpmTagVal = C.RPMTAG_PAYLOADFORMAT
	RPMTAG_PAYLOADCOMPRESSOR RpmTagVal = C.RPMTAG_PAYLOADCOMPRESSOR
	RPMTAG_PAYLOADFLAGS      RpmTagVal = C.RPMTAG_PAYLOADFLAGS
	RPMTAG_INSTALLCOLOR      RpmTagVal = C.RPMTAG_INSTALLCOLOR
	RPMTAG_INSTALLTID        RpmTagVal = C.RPMTAG_INSTALLTID
	RPMTAG_REMOVETID         RpmTagVal = C.RPMTAG_REMOVETID
	RPMTAG_SHA1RHN           RpmTagVal = C.RPMTAG_SHA1RHN
	RPMTAG_RHNPLATFORM       RpmTagVal = C.RPMTAG_RHNPLATFORM
	RPMTAG_PLATFORM          RpmTagVal = C.RPMTAG_PLATFORM
	RPMTAG_PATCHESNAME       RpmTagVal = C.RPMTAG_PATCHESNAME
	RPMTAG_PATCHESFLAGS      RpmTagVal = C.RPMTAG_PATCHESFLAGS
	RPMTAG_PATCHESVERSION    RpmTagVal = C.RPMTAG_PATCHESVERSION
	RPMTAG_CACHECTIME        RpmTagVal = C.RPMTAG_CACHECTIME
	RPMTAG_CACHEPKGPATH      RpmTagVal = C.RPMTAG_CACHEPKGPATH
	RPMTAG_CACHEPKGSIZE      RpmTagVal = C.RPMTAG_CACHEPKGSIZE
	RPMTAG_CACHEPKGMTIME     RpmTagVal = C.RPMTAG_CACHEPKGMTIME
	RPMTAG_FILECOLORS        RpmTagVal = C.RPMTAG_FILECOLORS
	RPMTAG_FILECLASS         RpmTagVal = C.RPMTAG_FILECLASS
	RPMTAG_CLASSDICT         RpmTagVal = C.RPMTAG_CLASSDICT
	RPMTAG_FILEDEPENDSX      RpmTagVal = C.RPMTAG_FILEDEPENDSX
	RPMTAG_FILEDEPENDSN      RpmTagVal = C.RPMTAG_FILEDEPENDSN
	RPMTAG_DEPENDSDICT       RpmTagVal = C.RPMTAG_DEPENDSDICT
	RPMTAG_SOURCEPKGID       RpmTagVal = C.RPMTAG_SOURCEPKGID
	RPMTAG_FILECONTEXTS      RpmTagVal = C.RPMTAG_FILECONTEXTS
	RPMTAG_FSCONTEXTS        RpmTagVal = C.RPMTAG_FSCONTEXTS
	RPMTAG_RECONTEXTS        RpmTagVal = C.RPMTAG_RECONTEXTS
	RPMTAG_POLICIES          RpmTagVal = C.RPMTAG_POLICIES
	RPMTAG_PRETRANS          RpmTagVal = C.RPMTAG_PRETRANS
	RPMTAG_POSTTRANS         RpmTagVal = C.RPMTAG_POSTTRANS
	RPMTAG_PRETRANSPROG      RpmTagVal = C.RPMTAG_PRETRANSPROG
	RPMTAG_POSTTRANSPROG     RpmTagVal = C.RPMTAG_POSTTRANSPROG
	RPMTAG_DISTTAG           RpmTagVal = C.RPMTAG_DISTTAG
	RPMTAG_SUGGESTS          RpmTagVal = C.RPMTAG_SUGGESTS
	RPMTAG_ENHANCES          RpmTagVal = C.RPMTAG_ENHANCES
	RPMTAG_PRIORITY          RpmTagVal = C.RPMTAG_PRIORITY
	RPMTAG_CVSID             RpmTagVal = C.RPMTAG_CVSID
	RPMTAG_SVNID             RpmTagVal = C.RPMTAG_SVNID
	RPMTAG_BLINKPKGID        RpmTagVal = C.RPMTAG_BLINKPKGID
	RPMTAG_BLINKHDRID        RpmTagVal = C.RPMTAG_BLINKHDRID
	RPMTAG_BLINKNEVRA        RpmTagVal = C.RPMTAG_BLINKNEVRA
	RPMTAG_FLINKPKGID        RpmTagVal = C.RPMTAG_FLINKPKGID
	RPMTAG_FLINKHDRID        RpmTagVal = C.RPMTAG_FLINKHDRID
	RPMTAG_FLINKNEVRA        RpmTagVal = C.RPMTAG_FLINKNEVRA
	RPMTAG_PACKAGEORIGIN     RpmTagVal = C.RPMTAG_PACKAGEORIGIN
	RPMTAG_TRIGGERPREIN      RpmTagVal = C.RPMTAG_TRIGGERPREIN
	RPMTAG_BUILDSUGGESTS     RpmTagVal = C.RPMTAG_BUILDSUGGESTS
	RPMTAG_BUILDENHANCES     RpmTagVal = C.RPMTAG_BUILDENHANCES
	RPMTAG_SCRIPTSTATES      RpmTagVal = C.RPMTAG_SCRIPTSTATES
	RPMTAG_SCRIPTMETRICS     RpmTagVal = C.RPMTAG_SCRIPTMETRICS
	RPMTAG_BUILDCPUCLOCK     RpmTagVal = C.RPMTAG_BUILDCPUCLOCK
	RPMTAG_FILEDIGESTALGOS   RpmTagVal = C.RPMTAG_FILEDIGESTALGOS
	RPMTAG_VARIANTS          RpmTagVal = C.RPMTAG_VARIANTS
	RPMTAG_XMAJOR            RpmTagVal = C.RPMTAG_XMAJOR
	RPMTAG_XMINOR            RpmTagVal = C.RPMTAG_XMINOR
	RPMTAG_REPOTAG           RpmTagVal = C.RPMTAG_REPOTAG
	RPMTAG_KEYWORDS          RpmTagVal = C.RPMTAG_KEYWORDS
	RPMTAG_BUILDPLATFORMS    RpmTagVal = C.RPMTAG_BUILDPLATFORMS
	RPMTAG_PACKAGECOLOR      RpmTagVal = C.RPMTAG_PACKAGECOLOR
	RPMTAG_PACKAGEPREFCOLOR  RpmTagVal = C.RPMTAG_PACKAGEPREFCOLOR
	RPMTAG_XATTRSDICT        RpmTagVal = C.RPMTAG_XATTRSDICT
	RPMTAG_FILEXATTRSX       RpmTagVal = C.RPMTAG_FILEXATTRSX
	RPMTAG_DEPATTRSDICT      RpmTagVal = C.RPMTAG_DEPATTRSDICT
	RPMTAG_CONFLICTATTRSX    RpmTagVal = C.RPMTAG_CONFLICTATTRSX
	RPMTAG_OBSOLETEATTRSX    RpmTagVal = C.RPMTAG_OBSOLETEATTRSX
	RPMTAG_PROVIDEATTRSX     RpmTagVal = C.RPMTAG_PROVIDEATTRSX
	RPMTAG_REQUIREATTRSX     RpmTagVal = C.RPMTAG_REQUIREATTRSX
	RPMTAG_BUILDPROVIDES     RpmTagVal = C.RPMTAG_BUILDPROVIDES
	RPMTAG_BUILDOBSOLETES    RpmTagVal = C.RPMTAG_BUILDOBSOLETES
	RPMTAG_DBINSTANCE        RpmTagVal = C.RPMTAG_DBINSTANCE
	RPMTAG_NVRA              RpmTagVal = C.RPMTAG_NVRA
	RPMTAG_FILENAMES         RpmTagVal = C.RPMTAG_FILENAMES
	RPMTAG_FILEPROVIDE       RpmTagVal = C.RPMTAG_FILEPROVIDE
	RPMTAG_FILEREQUIRE       RpmTagVal = C.RPMTAG_FILEREQUIRE
	RPMTAG_FSNAMES           RpmTagVal = C.RPMTAG_FSNAMES
	RPMTAG_FSSIZES           RpmTagVal = C.RPMTAG_FSSIZES
	RPMTAG_TRIGGERCONDS      RpmTagVal = C.RPMTAG_TRIGGERCONDS
	RPMTAG_TRIGGERTYPE       RpmTagVal = C.RPMTAG_TRIGGERTYPE
	RPMTAG_ORIGFILENAMES     RpmTagVal = C.RPMTAG_ORIGFILENAMES
	RPMTAG_LONGFILESIZES     RpmTagVal = C.RPMTAG_LONGFILESIZES
	RPMTAG_LONGSIZE          RpmTagVal = C.RPMTAG_LONGSIZE
	RPMTAG_FILECAPS          RpmTagVal = C.RPMTAG_FILECAPS
	RPMTAG_FILEDIGESTALGO    RpmTagVal = C.RPMTAG_FILEDIGESTALGO
	RPMTAG_BUGURL            RpmTagVal = C.RPMTAG_BUGURL
	RPMTAG_EVR               RpmTagVal = C.RPMTAG_EVR
	RPMTAG_NVR               RpmTagVal = C.RPMTAG_NVR
	RPMTAG_NEVR              RpmTagVal = C.RPMTAG_NEVR
	RPMTAG_NEVRA             RpmTagVal = C.RPMTAG_NEVRA
	RPMTAG_HEADERCOLOR       RpmTagVal = C.RPMTAG_HEADERCOLOR
	RPMTAG_VERBOSE           RpmTagVal = C.RPMTAG_VERBOSE
	RPMTAG_EPOCHNUM          RpmTagVal = C.RPMTAG_EPOCHNUM
	RPMTAG_ORDERNAME         RpmTagVal = C.RPMTAG_ORDERNAME
	RPMTAG_ORDERVERSION      RpmTagVal = C.RPMTAG_ORDERVERSION
	RPMTAG_ORDERFLAGS        RpmTagVal = C.RPMTAG_ORDERFLAGS
	RPMTAG_FIRSTFREE_TAG     RpmTagVal = C.RPMTAG_FIRSTFREE_TAG
)

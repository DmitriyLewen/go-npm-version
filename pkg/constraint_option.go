package npm

type conf struct {
	includePreRelease  bool
	allowBuildMetadata bool
}

type ConstraintOption interface {
	apply(*conf)
}

type WithPreRelease bool

func (o WithPreRelease) apply(c *conf) {
	c.includePreRelease = bool(o)
}

// AllowBuildMetadata is an option that makes build metadata significant.
// Semantic Versioning requires build metadata to be ignored and defines no
// ordering for it, so by default "1.2.3+build.1" and "1.2.3+build.2" are equal.
// This option is a deliberate extension that applies the pre-release precedence
// rules to build metadata: identifiers are compared one by one, numeric ones are
// compared numerically and rank below alphanumeric ones, and a longer set of
// identifiers ranks higher:
// "1.2.3" < "1.2.3+build" < "1.2.3+build.1" < "1.2.3+build.9" <
// "1.2.3+build.10" < "1.2.3+build.alpha".
// Other Semantic Versioning implementations, npm included, don't reproduce this
// order.
// Constraints with a wildcard (e.g. "2", "1.2.x" or "1.2.3-x") still ignore
// build metadata.
//
// The option applies to every comparison, including constraints that carry no
// metadata: with it enabled "1.2.3+build.1" no longer satisfies "=1.2.3" or
// "<=1.2.3", and does satisfy ">1.2.3".
// It affects constraint checking only: Version comparison (LessThan, Equal,
// Compare) always follows Semantic Versioning.
type AllowBuildMetadata bool

func (o AllowBuildMetadata) apply(c *conf) {
	c.allowBuildMetadata = bool(o)
}

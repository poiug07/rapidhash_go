//go:build !rapidhash_protected

package rapidhash

/*
 *  rapidhash configuration (default build).

 * 	To use protected build/test with:
 *      go build -tags rapidhash_protected
 *      go test -tags rapidhash_protected
 */

/*
 *  Protection option, alters behaviour of the rapid_mum multiplication
 *  function used by every variant.
 *
 *  protected = false (default): "fast" mode, max speed.
 */
const protected = false

//go:nosplit
func rapid_mix(a, b uint64) uint64 {
	hi, lo := rapid_mum(a, b)
	return lo ^ hi
}

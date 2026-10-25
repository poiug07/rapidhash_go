//go:build rapidhash_protected

package rapidhash

/*
 *  Protection option, alters behaviour of the rapid_mum multiplication
 *  function used by every variant.
 *
 *  protected = true: extra protection against entropy loss.
 */
const protected = true

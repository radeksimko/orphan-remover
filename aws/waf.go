package aws

import (
	"log"

	awsSDK "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/waf"
)

func (a *AWS) RemoveWaf() error {
	err := a.RemoveWafXssMatchSets()
	if err != nil {
		return err
	}

	err = a.RemoveWafIPSets()
	if err != nil {
		return err
	}

	err = a.RemoveWafSizeConstraintSets()
	if err != nil {
		return err
	}

	err = a.RemoveWafByteMatchSets()
	if err != nil {
		return err
	}

	return nil
}

func (a *AWS) RemoveWafXssMatchSets() error {
	out, err := a.wafConn.ListXssMatchSets(&waf.ListXssMatchSetsInput{})
	if err != nil {
		return err
	}

	log.Printf("[INFO] Found %d XSS match sets", len(out.XssMatchSets))

	// TODO: Pagination
	for _, s := range out.XssMatchSets {
		set, err := a.wafConn.GetXssMatchSet(&waf.GetXssMatchSetInput{
			XssMatchSetId: s.XssMatchSetId,
		})
		if err != nil {
			return err
		}

		for _, mt := range set.XssMatchSet.XssMatchTuples {
			out, err := a.wafConn.GetChangeToken(&waf.GetChangeTokenInput{})
			if err != nil {
				return err
			}
			input := waf.UpdateXssMatchSetInput{
				ChangeToken:   out.ChangeToken,
				XssMatchSetId: s.XssMatchSetId,
				Updates: []*waf.XssMatchSetUpdate{
					{
						Action:        awsSDK.String(waf.ChangeActionDelete),
						XssMatchTuple: mt,
					},
				},
			}
			_, err = a.wafConn.UpdateXssMatchSet(&input)
			if err != nil {
				return err
			}
			log.Printf("[INFO] Removed XSS match set tuple: %s", mt.String())
		}

		out, err := a.wafConn.GetChangeToken(&waf.GetChangeTokenInput{})
		if err != nil {
			return err
		}

		_, err = a.wafConn.DeleteXssMatchSet(&waf.DeleteXssMatchSetInput{
			ChangeToken:   out.ChangeToken,
			XssMatchSetId: s.XssMatchSetId,
		})
		if err != nil {
			return err
		}
		log.Printf("[INFO] Deleted XSS match set: %q", *s.XssMatchSetId)
	}
	return nil
}

func (a *AWS) RemoveWafIPSets() error {
	out, err := a.wafConn.ListIPSets(&waf.ListIPSetsInput{})
	if err != nil {
		return err
	}

	log.Printf("[INFO] Found %d IP sets", len(out.IPSets))

	// TODO: Pagination
	for _, s := range out.IPSets {
		set, err := a.wafConn.GetIPSet(&waf.GetIPSetInput{
			IPSetId: s.IPSetId,
		})
		if err != nil {
			return err
		}

		for _, sd := range set.IPSet.IPSetDescriptors {
			out, err := a.wafConn.GetChangeToken(&waf.GetChangeTokenInput{})
			if err != nil {
				return err
			}
			input := waf.UpdateIPSetInput{
				ChangeToken: out.ChangeToken,
				IPSetId:     s.IPSetId,
				Updates: []*waf.IPSetUpdate{
					{
						Action:          awsSDK.String(waf.ChangeActionDelete),
						IPSetDescriptor: sd,
					},
				},
			}
			_, err = a.wafConn.UpdateIPSet(&input)
			if err != nil {
				return err
			}
			log.Printf("[INFO] Removed IP set descriptor: %s", sd.String())
		}

		out, err := a.wafConn.GetChangeToken(&waf.GetChangeTokenInput{})
		if err != nil {
			return err
		}

		_, err = a.wafConn.DeleteIPSet(&waf.DeleteIPSetInput{
			ChangeToken: out.ChangeToken,
			IPSetId:     s.IPSetId,
		})
		if err != nil {
			return err
		}
		log.Printf("[INFO] Deleted IP set: %q", *s.IPSetId)
	}
	return nil
}

func (a *AWS) RemoveWafSizeConstraintSets() error {
	out, err := a.wafConn.ListSizeConstraintSets(&waf.ListSizeConstraintSetsInput{})
	if err != nil {
		return err
	}

	log.Printf("[INFO] Found %d Size Constraint sets", len(out.SizeConstraintSets))

	// TODO: Pagination
	for _, s := range out.SizeConstraintSets {
		set, err := a.wafConn.GetSizeConstraintSet(&waf.GetSizeConstraintSetInput{
			SizeConstraintSetId: s.SizeConstraintSetId,
		})
		if err != nil {
			return err
		}

		for _, sc := range set.SizeConstraintSet.SizeConstraints {
			out, err := a.wafConn.GetChangeToken(&waf.GetChangeTokenInput{})
			if err != nil {
				return err
			}
			input := waf.UpdateSizeConstraintSetInput{
				ChangeToken:         out.ChangeToken,
				SizeConstraintSetId: s.SizeConstraintSetId,
				Updates: []*waf.SizeConstraintSetUpdate{
					{
						Action:         awsSDK.String(waf.ChangeActionDelete),
						SizeConstraint: sc,
					},
				},
			}
			_, err = a.wafConn.UpdateSizeConstraintSet(&input)
			if err != nil {
				return err
			}
			log.Printf("[INFO] Removed Size Constraint: %s", sc.String())
		}

		out, err := a.wafConn.GetChangeToken(&waf.GetChangeTokenInput{})
		if err != nil {
			return err
		}

		_, err = a.wafConn.DeleteSizeConstraintSet(&waf.DeleteSizeConstraintSetInput{
			ChangeToken:         out.ChangeToken,
			SizeConstraintSetId: s.SizeConstraintSetId,
		})
		if err != nil {
			return err
		}
		log.Printf("[INFO] Deleted Size Constraint set: %q", *s.SizeConstraintSetId)
	}
	return nil
}

func (a *AWS) RemoveWafByteMatchSets() error {
	out, err := a.wafConn.ListByteMatchSets(&waf.ListByteMatchSetsInput{})
	if err != nil {
		return err
	}

	log.Printf("[INFO] Found %d Byte Match sets", len(out.ByteMatchSets))

	// TODO: Pagination
	for _, s := range out.ByteMatchSets {
		set, err := a.wafConn.GetByteMatchSet(&waf.GetByteMatchSetInput{
			ByteMatchSetId: s.ByteMatchSetId,
		})
		if err != nil {
			return err
		}

		for _, t := range set.ByteMatchSet.ByteMatchTuples {
			out, err := a.wafConn.GetChangeToken(&waf.GetChangeTokenInput{})
			if err != nil {
				return err
			}
			input := waf.UpdateByteMatchSetInput{
				ChangeToken:    out.ChangeToken,
				ByteMatchSetId: s.ByteMatchSetId,
				Updates: []*waf.ByteMatchSetUpdate{
					{
						Action:         awsSDK.String(waf.ChangeActionDelete),
						ByteMatchTuple: t,
					},
				},
			}
			_, err = a.wafConn.UpdateByteMatchSet(&input)
			if err != nil {
				return err
			}
			log.Printf("[INFO] Removed Byte Match tuple: %s", t.String())
		}

		out, err := a.wafConn.GetChangeToken(&waf.GetChangeTokenInput{})
		if err != nil {
			return err
		}

		_, err = a.wafConn.DeleteByteMatchSet(&waf.DeleteByteMatchSetInput{
			ChangeToken:    out.ChangeToken,
			ByteMatchSetId: s.ByteMatchSetId,
		})
		if err != nil {
			return err
		}
		log.Printf("[INFO] Deleted Byte Match set: %q", *s.ByteMatchSetId)
	}
	return nil
}

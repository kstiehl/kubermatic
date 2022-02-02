/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apiserver

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/url"
	"sort"
	"time"

	"k8c.io/kubermatic/v2/pkg/controller/operator/common"
	"k8c.io/kubermatic/v2/pkg/controller/operator/seed/resources/nodeportproxy"
	kubermaticv1 "k8c.io/kubermatic/v2/pkg/crd/kubermatic/v1"
	kubermaticmaster "k8c.io/kubermatic/v2/pkg/install/stack/kubermatic-master"
	"k8c.io/kubermatic/v2/pkg/resources"
	"k8c.io/kubermatic/v2/pkg/resources/reconciling"

	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DenyAllPolicyCreator returns a func to create/update the apiserver
// deny all egress policy.
func DenyAllPolicyCreator() reconciling.NamedNetworkPolicyCreatorGetter {
	return func() (string, reconciling.NetworkPolicyCreator) {
		return resources.NetworkPolicyDefaultDenyAllEgress, func(np *networkingv1.NetworkPolicy) (*networkingv1.NetworkPolicy, error) {
			np.Spec = networkingv1.NetworkPolicySpec{
				PolicyTypes: []networkingv1.PolicyType{
					networkingv1.PolicyTypeEgress,
				},
				PodSelector: metav1.LabelSelector{
					MatchLabels: map[string]string{
						resources.AppLabelKey: name,
					},
				},
				Egress: []networkingv1.NetworkPolicyEgressRule{},
			}

			return np, nil
		}
	}
}

// EctdAllowCreator returns a func to create/update the apiserver ETCD allow egress policy.
func EctdAllowCreator(c *kubermaticv1.Cluster) reconciling.NamedNetworkPolicyCreatorGetter {
	return func() (string, reconciling.NetworkPolicyCreator) {
		return resources.NetworkPolicyEtcdAllow, func(np *networkingv1.NetworkPolicy) (*networkingv1.NetworkPolicy, error) {
			np.Spec = networkingv1.NetworkPolicySpec{
				PolicyTypes: []networkingv1.PolicyType{
					networkingv1.PolicyTypeEgress,
				},
				PodSelector: metav1.LabelSelector{
					MatchLabels: map[string]string{
						resources.AppLabelKey: name,
					},
				},
				Egress: []networkingv1.NetworkPolicyEgressRule{
					{
						To: []networkingv1.NetworkPolicyPeer{
							{
								PodSelector: &metav1.LabelSelector{
									MatchLabels: map[string]string{
										resources.AppLabelKey: "etcd",
										"cluster":             c.ObjectMeta.Name,
									},
								},
							},
						},
					},
				},
			}

			return np, nil
		}
	}
}

// DNSAllowCreator returns a func to create/update the apiserver DNS allow egress policy.
func DNSAllowCreator(c *kubermaticv1.Cluster, data *resources.TemplateData) reconciling.NamedNetworkPolicyCreatorGetter {
	return func() (string, reconciling.NetworkPolicyCreator) {
		return resources.NetworkPolicyDNSAllow, func(np *networkingv1.NetworkPolicy) (*networkingv1.NetworkPolicy, error) {
			dnsPort := intstr.FromInt(53)
			protoUdp := corev1.ProtocolUDP
			protoTcp := corev1.ProtocolTCP

			np.Spec = networkingv1.NetworkPolicySpec{
				PolicyTypes: []networkingv1.PolicyType{
					networkingv1.PolicyTypeEgress,
				},
				PodSelector: metav1.LabelSelector{
					MatchLabels: map[string]string{
						resources.AppLabelKey: name,
					},
				},
				Egress: []networkingv1.NetworkPolicyEgressRule{
					{
						Ports: []networkingv1.NetworkPolicyPort{
							{
								Protocol: &protoUdp,
								Port:     &dnsPort,
							},
							{
								Protocol: &protoTcp,
								Port:     &dnsPort,
							},
						},
					},
				},
			}
			return np, nil
		}
	}
}

// OpenVPNServerAllowCreator returns a func to create/update the apiserver OpenVPN allow egress policy.
func OpenVPNServerAllowCreator(c *kubermaticv1.Cluster) reconciling.NamedNetworkPolicyCreatorGetter {
	return func() (string, reconciling.NetworkPolicyCreator) {
		return resources.NetworkPolicyOpenVPNServerAllow, func(np *networkingv1.NetworkPolicy) (*networkingv1.NetworkPolicy, error) {
			np.Spec = networkingv1.NetworkPolicySpec{
				PolicyTypes: []networkingv1.PolicyType{
					networkingv1.PolicyTypeEgress,
				},
				PodSelector: metav1.LabelSelector{
					MatchLabels: map[string]string{
						resources.AppLabelKey: name,
					},
				},
				Egress: []networkingv1.NetworkPolicyEgressRule{
					{
						To: []networkingv1.NetworkPolicyPeer{
							{
								PodSelector: &metav1.LabelSelector{
									MatchLabels: map[string]string{
										resources.AppLabelKey: "openvpn-server",
										"cluster":             c.ObjectMeta.Name,
									},
								},
							},
						},
					},
				},
			}

			return np, nil
		}
	}
}

func MachineControllerWebhookCreator(c *kubermaticv1.Cluster) reconciling.NamedNetworkPolicyCreatorGetter {
	return func() (string, reconciling.NetworkPolicyCreator) {
		return resources.NetworkPolicyMachineControllerWebhookAllow, func(np *networkingv1.NetworkPolicy) (*networkingv1.NetworkPolicy, error) {
			np.Spec = networkingv1.NetworkPolicySpec{
				PolicyTypes: []networkingv1.PolicyType{
					networkingv1.PolicyTypeEgress,
				},
				PodSelector: metav1.LabelSelector{
					MatchLabels: map[string]string{
						resources.AppLabelKey: name,
					},
				},
				Egress: []networkingv1.NetworkPolicyEgressRule{
					{
						To: []networkingv1.NetworkPolicyPeer{
							{
								PodSelector: &metav1.LabelSelector{
									MatchLabels: map[string]string{
										resources.AppLabelKey: "machine-controller-webhook",
									},
								},
							},
						},
					},
				},
			}

			return np, nil
		}
	}
}

func MetricsServerAllowCreator(c *kubermaticv1.Cluster) reconciling.NamedNetworkPolicyCreatorGetter {
	return func() (string, reconciling.NetworkPolicyCreator) {
		return resources.NetworkPolicyMetricsServerAllow, func(np *networkingv1.NetworkPolicy) (*networkingv1.NetworkPolicy, error) {
			np.Spec = networkingv1.NetworkPolicySpec{
				PolicyTypes: []networkingv1.PolicyType{
					networkingv1.PolicyTypeEgress,
				},
				PodSelector: metav1.LabelSelector{
					MatchLabels: map[string]string{
						resources.AppLabelKey: name,
					},
				},
				Egress: []networkingv1.NetworkPolicyEgressRule{
					{
						To: []networkingv1.NetworkPolicyPeer{
							{
								PodSelector: &metav1.LabelSelector{
									MatchLabels: map[string]string{
										resources.AppLabelKey: "metrics-server",
									},
								},
							},
						},
					},
				},
			}

			return np, nil
		}
	}
}

// ClusterExternalAddrAllowCreator returns a func to create/update the apiserver cluster-external-addr-allow egress policy.
// This policy is necessary in Konnectivity setup, so that konnectivity-server can connect to the apiserver via
// the external URL (used as service-account-issuer) to validate konnectivity-agent authentication token.
func ClusterExternalAddrAllowCreator(c *kubermaticv1.Cluster) reconciling.NamedNetworkPolicyCreatorGetter {
	return func() (string, reconciling.NetworkPolicyCreator) {
		return resources.NetworkPolicyClusterExternalAddrAllow, func(np *networkingv1.NetworkPolicy) (*networkingv1.NetworkPolicy, error) {
			np.Spec = networkingv1.NetworkPolicySpec{
				PolicyTypes: []networkingv1.PolicyType{
					networkingv1.PolicyTypeEgress,
				},
				PodSelector: metav1.LabelSelector{
					MatchLabels: map[string]string{
						resources.AppLabelKey: name,
					},
				},
				Egress: []networkingv1.NetworkPolicyEgressRule{
					{
						To: []networkingv1.NetworkPolicyPeer{
							{
								// allow egress traffic to the nodeport-proxy as for some CNI + kube-proxy mode
								// combinations a local path to it may be used to reach the external apiserver address
								NamespaceSelector: &metav1.LabelSelector{
									MatchLabels: map[string]string{
										common.NameLabel: nodeportproxy.EnvoyDeploymentName,
									},
								},
							},
						},
					},
				},
			}
			// allow egress traffic to all resolved cluster external IPs
			ipList, err := hostnameToIPList(c.Address.ExternalName)
			if err != nil {
				return nil, fmt.Errorf("failed to resolve cluster external name %s: %w", c.Address.ExternalName, err)
			}
			for _, ip := range ipList {
				cidr := fmt.Sprintf("%s/%d", ip.String(), net.IPv4len*8)
				if ip.To4() == nil {
					cidr = fmt.Sprintf("%s/%d", ip.String(), net.IPv6len*8)
				}
				np.Spec.Egress[0].To = append(np.Spec.Egress[0].To, networkingv1.NetworkPolicyPeer{
					IPBlock: &networkingv1.IPBlock{
						CIDR: cidr,
					},
				})
			}
			return np, nil
		}
	}
}

// OIDCIssuerAllowCreator returns a func to create/update the apiserver oidc-issuer-allow egress policy.
func OIDCIssuerAllowCreator(issuerURL string) reconciling.NamedNetworkPolicyCreatorGetter {
	return func() (string, reconciling.NetworkPolicyCreator) {
		return resources.NetworkPolicyOIDCIssuerAllow, func(np *networkingv1.NetworkPolicy) (*networkingv1.NetworkPolicy, error) {
			np.Spec = networkingv1.NetworkPolicySpec{
				PolicyTypes: []networkingv1.PolicyType{
					networkingv1.PolicyTypeEgress,
				},
				PodSelector: metav1.LabelSelector{
					MatchLabels: map[string]string{
						resources.AppLabelKey: name,
					},
				},
				Egress: []networkingv1.NetworkPolicyEgressRule{
					{
						To: []networkingv1.NetworkPolicyPeer{
							{
								// allow egress traffic to the nginx-ingress-controller as for some CNI + kube-proxy
								// mode combinations a local path to it may be used to reach OIDC issuer installed in KKP
								NamespaceSelector: &metav1.LabelSelector{
									MatchLabels: map[string]string{
										corev1.LabelMetadataName: kubermaticmaster.NginxIngressControllerNamespace,
									},
								},
							},
						},
					},
				},
			}
			// allow egress traffic to OIDC issuer's external IPs
			u, err := url.Parse(issuerURL)
			if err != nil {
				return nil, fmt.Errorf("failed to parse OIDC issuer URL %s: %w", issuerURL, err)
			}
			ipList, err := hostnameToIPList(u.Hostname())
			if err != nil {
				return nil, fmt.Errorf("failed to resolve OIDC issuer hostname %s: %w", u.Hostname(), err)
			}
			for _, ip := range ipList {
				cidr := fmt.Sprintf("%s/%d", ip.String(), net.IPv4len*8)
				if ip.To4() == nil {
					cidr = fmt.Sprintf("%s/%d", ip.String(), net.IPv6len*8)
				}
				np.Spec.Egress[0].To = append(np.Spec.Egress[0].To, networkingv1.NetworkPolicyPeer{
					IPBlock: &networkingv1.IPBlock{
						CIDR: cidr,
					},
				})
			}
			return np, nil
		}
	}
}

// hostnameToIPList returns a list of IP addresses used to reach the provided hostname.
// If it is an IP address, returns it. If it is a domain name, resolves it.
// The returned list of IPs is always sorted to produce the same result on each resolution attempt.
func hostnameToIPList(hostname string) ([]net.IP, error) {
	if ip := net.ParseIP(hostname); ip != nil {
		// hostname is an IP address
		return []net.IP{ip}, nil
	}
	// hostname is a domain name - resolve it
	ipList, err := lookupIPWithTimeout(hostname, 5*time.Second)
	if err != nil {
		return nil, err
	}
	if len(ipList) == 0 {
		return nil, fmt.Errorf("no resolved IP address for hostname %s", hostname)
	}
	sort.Slice(ipList, func(i, j int) bool {
		return bytes.Compare(ipList[i], ipList[j]) < 0
	})
	return ipList, nil
}

func lookupIPWithTimeout(host string, timeout time.Duration) ([]net.IP, error) {
	var r net.Resolver
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel() // to avoid possible resource leak
	return r.LookupIP(ctx, "ip", host)
}

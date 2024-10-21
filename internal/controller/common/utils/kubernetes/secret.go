package common

import (
	password "github.com/sethvargo/go-password/password"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Create a password secret
func generateSecretHash(password string) string {
	return password
}

func createSecret(name string, namespace string) *corev1.Secret {
	random, err := password.Generate(43, 10, 0, false, false)
	if err != nil {
		// Fail kindly and generate a crappy password
		random = "defaultPassword"
	}
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: map[string][]byte{
			"password": []byte(random),
		},
	}
	return secret
}

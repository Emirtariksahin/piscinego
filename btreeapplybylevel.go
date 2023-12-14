package piscine

// BTreeApplyByLevel, bir ağaç düğümünün kökünden başlayarak seviye seviye dolaşıp
// her düğümün verisine belirtilen fonksiyonu uygular.
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	// Kök düğüm nil ise fonksiyonu sonlandır
	if root == nil {
		return
	}

	// Bir kuyruk (queue) oluştur ve kök düğümü kuyruğa ekle
	queue := []*TreeNode{root}

	// Kuyruk boş olana kadar döngüyü devam ettir
	for len(queue) > 0 {
		// Kuyruktan bir düğümü çıkar
		currentNode := queue[0]
		queue = queue[1:]

		// Belirtilen fonksiyonu, mevcut düğümün verisi üzerinde uygula
		_, _ = f(currentNode.Data)

		// Sol çocuk düğümü varsa kuyruğa ekle
		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}

		// Sağ çocuk düğümü varsa kuyruğa ekle
		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}
	}
}

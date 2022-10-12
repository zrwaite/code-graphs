//
//  TextViewModifier.swift
//  CodeGraphs
//
//  Created by Zachary Waite on 2022-10-11.
//

import Foundation
import SwiftUI

struct WWText: ViewModifier {
    func body(content: Content) -> some View {
        content
            .foregroundColor(Color.white)
            .padding()
            .multilineTextAlignment(.center)
    }
}
